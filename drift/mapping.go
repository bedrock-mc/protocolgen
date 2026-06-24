package main

import (
	"fmt"
	"slices"
)

// Comparison statuses between a documented field and a source wire op.
const (
	statusOK       = "ok"       // encoding agrees with the docs
	statusMismatch = "mismatch" // encoding disagrees (likely real drift)
	statusReview   = "review"   // structural/handcoded op — a human must judge
)

// knownComposites maps Mojang composite definition titles to the gophertunnel IO
// method that encodes them as a single op.
var knownComposites = map[string]string{
	"BlockPos":             "BlockPos",
	"NetworkBlockPosition": "BlockPos",
	"Vec3":                 "Vec3",
	"Vec2":                 "Vec2",
	"ChunkPos":             "ChunkPos",
	"SubChunkPos":          "SubChunkPos",
}

// compatibleEncoding reports whether a source wire op is an acceptable encoding
// for a documented field, returning a status and an explanatory note.
func compatibleEncoding(doc DocField, op WireOp) (status, note string) {
	// Only direct io.Method(...) ops can be checked automatically. protocol
	// helpers (Slice, OptionalFunc, ...) and custom handcoded calls are
	// structural and need a human to judge.
	if op.Kind != "io" {
		return statusReview, fmt.Sprintf("%s op %q is structural/handcoded — verify against doc field manually", op.Kind, op.Method)
	}

	if doc.Composite {
		want, ok := knownComposites[doc.RefTitle]
		if !ok {
			return statusReview, fmt.Sprintf("composite %q has no known single-op mapping", doc.RefTitle)
		}
		if op.Method == want {
			return statusOK, ""
		}
		return statusMismatch, fmt.Sprintf("doc composite %q expects io.%s, source uses io.%s", doc.RefTitle, want, op.Method)
	}

	accepted := acceptedIOMethods(doc)
	if len(accepted) == 0 {
		return statusReview, fmt.Sprintf("no encoding mapping for underlying-type %q", doc.UnderlyingType)
	}
	if slices.Contains(accepted, op.Method) {
		return statusOK, ""
	}
	return statusMismatch, fmt.Sprintf("doc %s expects io.%v, source uses io.%s", docTypeDesc(doc), accepted, op.Method)
}

// acceptedIOMethods returns the gophertunnel IO methods that correctly encode a
// documented scalar field.
func acceptedIOMethods(doc DocField) []string {
	// An enum without "Enum-as-Value" is serialized as its string name.
	if doc.Enum && !doc.HasOption("Enum-as-Value") {
		return []string{"String", "StringUTF"}
	}
	comp := doc.HasOption("Compression")
	switch doc.UnderlyingType {
	case "bool", "boolean":
		return []string{"Bool"}
	case "int8":
		return []string{"Int8"}
	case "uint8":
		return []string{"Uint8"}
	case "int16":
		return []string{"Int16"}
	case "uint16":
		return []string{"Uint16"}
	case "int32":
		if comp {
			return []string{"Varint32"}
		}
		return []string{"Int32", "BEInt32"}
	case "uint32":
		if comp {
			return []string{"Varuint32"}
		}
		return []string{"Uint32"}
	case "int64":
		if comp {
			return []string{"Varint64"}
		}
		return []string{"Int64"}
	case "uint64":
		if comp {
			return []string{"Varuint64"}
		}
		return []string{"Uint64"}
	case "float":
		return []string{"Float32", "ByteFloat"}
	case "double":
		return []string{"Float64"}
	case "string":
		// Mojang types both real strings and NBT/byte blobs as "string".
		return []string{"String", "StringUTF", "NBT", "ByteSlice", "Bytes"}
	default:
		return nil
	}
}

// docTypeDesc describes a documented scalar field for diagnostics.
func docTypeDesc(doc DocField) string {
	if doc.HasOption("Compression") {
		return doc.UnderlyingType + "+Compression"
	}
	return doc.UnderlyingType
}
