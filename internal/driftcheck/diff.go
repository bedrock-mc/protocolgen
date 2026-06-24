package driftcheck

import "fmt"

// Field-level statuses produced by alignment (in addition to the encoding
// statuses in mapping.go).
const (
	statusMissingSrc = "missing-in-source" // doc describes a field the source does not serialize
	statusExtraSrc   = "extra-in-source"   // source serializes an op the docs do not describe
)

// FieldDiff is the comparison of one documented field against one source wire op.
type FieldDiff struct {
	Ordinal  int
	DocName  string
	DocType  string
	SrcField string
	SrcOp    string
	Guarded  bool
	Status   string
	Note     string
}

// PacketDiff is the per-packet comparison result.
type PacketDiff struct {
	ID      int
	DocName string
	SrcName string
	Fields  []FieldDiff
}

// Report is the full comparison between docs and source.
type Report struct {
	MissingInSource []DocPacket    // present in docs, absent in source
	MissingInDocs   []SourcePacket // present in source, absent in docs
	Diffs           []PacketDiff
}

// Compare aligns documented packets against source packets by numeric ID.
func Compare(docs []DocPacket, srcs []SourcePacket) Report {
	srcByID := make(map[int]SourcePacket, len(srcs))
	for _, s := range srcs {
		srcByID[s.ID] = s
	}
	docByID := make(map[int]bool, len(docs))

	var rep Report
	for _, doc := range docs {
		docByID[doc.ID] = true
		src, ok := srcByID[doc.ID]
		if !ok {
			rep.MissingInSource = append(rep.MissingInSource, doc)
			continue
		}
		rep.Diffs = append(rep.Diffs, diffPacket(doc, src))
	}
	for _, src := range srcs {
		if !docByID[src.ID] {
			rep.MissingInDocs = append(rep.MissingInDocs, src)
		}
	}
	return rep
}

// diffPacket aligns a doc and source packet and classifies each pairing. It is a
// two-pointer walk so a documented composite (e.g. Vec2) can match either one
// dedicated source op (io.Vec2) or its constituent scalar ops (two Float32),
// which avoids a cascade of false mismatches when the two sides group fields
// differently.
func diffPacket(doc DocPacket, src SourcePacket) PacketDiff {
	d := PacketDiff{ID: doc.ID, DocName: doc.Name, SrcName: src.TypeName}

	// A packet with conditional ops cannot be aligned reliably by position, so
	// field-count divergences are downgraded to review rather than asserted as
	// drift. Encoding mismatches are still reported regardless.
	hasGuard := false
	for _, op := range src.Ops {
		if op.Guarded {
			hasGuard = true
			break
		}
	}

	fields, ops := doc.Fields, src.Ops
	di, si, ord := 0, 0, 0
	for di < len(fields) || si < len(ops) {
		switch {
		case di < len(fields) && si < len(ops):
			field, op := fields[di], ops[si]
			// A composite gophertunnel expanded into scalar ops: match its
			// constituent scalars against the next ops.
			if field.Composite && op.Kind == "io" {
				if _, typed := knownComposites[field.RefTitle]; !typed || op.Method != knownComposites[field.RefTitle] {
					if cons := compositeConstituents(field.RefTitle); len(cons) > 0 && si+len(cons) <= len(ops) && scalarsMatch(cons, ops[si:si+len(cons)]) {
						d.Fields = append(d.Fields, FieldDiff{
							Ordinal: ord, DocName: field.Name, DocType: fieldDesc(field),
							SrcField: ops[si].Field, SrcOp: fmt.Sprintf("%d× io.%s (expanded)", len(cons), ops[si].Method),
							Status: statusOK, Note: "composite expanded into scalar ops",
						})
						di++
						si += len(cons)
						ord++
						continue
					}
				}
			}
			st, note := compatibleEncoding(field, op)
			d.Fields = append(d.Fields, FieldDiff{
				Ordinal: ord, DocName: field.Name, DocType: fieldDesc(field),
				SrcField: op.Field, SrcOp: op.Kind + "." + op.Method, Guarded: op.Guarded,
				Status: st, Note: note,
			})
			di++
			si++
			ord++
		case di < len(fields):
			field := fields[di]
			fd := FieldDiff{Ordinal: ord, DocName: field.Name, DocType: fieldDesc(field), SrcField: "—", SrcOp: "—"}
			if hasGuard {
				fd.Status = statusReview
				fd.Note = "fewer ops than documented fields in a conditional packet — verify manually"
			} else {
				fd.Status = statusMissingSrc
				fd.Note = "source has no wire op at this position"
			}
			d.Fields = append(d.Fields, fd)
			di++
			ord++
		default:
			op := ops[si]
			fd := FieldDiff{Ordinal: ord, DocName: "—", DocType: "—", SrcField: op.Field, SrcOp: op.Kind + "." + op.Method, Guarded: op.Guarded}
			if op.Guarded || hasGuard {
				fd.Status = statusReview
				fd.Note = "op beyond documented fields in a conditional packet (likely optional/union handling)"
			} else {
				fd.Status = statusExtraSrc
				fd.Note = "docs describe no field at this position"
			}
			d.Fields = append(d.Fields, fd)
			si++
			ord++
		}
	}
	return d
}

// compositeConstituents returns the scalar fields a known composite decomposes
// into on the wire, for matching against expanded source ops.
func compositeConstituents(title string) []DocField {
	flt := DocField{UnderlyingType: "float", JSONType: "number"}
	i32c := DocField{UnderlyingType: "int32", Options: []string{"Compression"}, JSONType: "integer"}
	switch title {
	case "Vec2":
		return []DocField{flt, flt}
	case "Vec3":
		return []DocField{flt, flt, flt}
	case "BlockPos", "NetworkBlockPosition", "SubChunkPos":
		return []DocField{i32c, i32c, i32c}
	case "ChunkPos":
		return []DocField{i32c, i32c}
	default:
		return nil
	}
}

// scalarsMatch reports whether each op cleanly encodes the corresponding scalar.
func scalarsMatch(cons []DocField, ops []WireOp) bool {
	for i := range cons {
		if st, _ := compatibleEncoding(cons[i], ops[i]); st != statusOK {
			return false
		}
	}
	return true
}

// fieldDesc renders a documented field's type for display.
func fieldDesc(f DocField) string {
	if f.Composite {
		return "composite " + f.RefTitle
	}
	if f.Enum && !f.HasOption("Enum-as-Value") {
		return "enum(string) " + f.UnderlyingType
	}
	return docTypeDesc(f)
}

// HasDrift reports whether the comparison found anything a maintainer must act
// on: an encoding mismatch, a field count divergence, or a packet add/remove.
// Review-only findings do not count as drift.
func (r Report) HasDrift() bool {
	if len(r.MissingInSource) > 0 || len(r.MissingInDocs) > 0 {
		return true
	}
	for _, d := range r.Diffs {
		for _, f := range d.Fields {
			switch f.Status {
			case statusMismatch, statusMissingSrc, statusExtraSrc:
				return true
			}
		}
	}
	return false
}
