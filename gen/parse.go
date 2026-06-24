package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// composite is a multi-field doc definition that the raw format generates as its
// own struct type.
type composite struct {
	Title  string
	GoName string
	Fields []genField
}

// collectComposites parses every doc in dir and returns each unique multi-field
// composite definition (single-scalar wrappers are flattened into their field,
// so they are excluded). Deduplicated by generated Go name.
func collectComposites(dir string) ([]composite, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read docs dir: %w", err)
	}
	seen := map[string]composite{}
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".json") || strings.HasPrefix(name, "__") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			continue
		}
		var doc rawGenDoc
		if json.Unmarshal(b, &doc) != nil {
			continue
		}
		for _, def := range doc.Definitions {
			if def.Title == "" || len(def.Properties) <= 1 {
				continue // single-scalar wrappers are flattened, not generated
			}
			gn := cleanCompositeTitle(def.Title)
			if gn == "" || seen[gn].GoName != "" {
				continue
			}
			var fields []genField
			for fname, rp := range def.Properties {
				fields = append(fields, buildField(fname, rp, doc.Definitions))
			}
			sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
			seen[gn] = composite{Title: def.Title, GoName: gn, Fields: fields}
		}
	}
	out := make([]composite, 0, len(seen))
	for _, c := range seen {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GoName < out[j].GoName })
	return out, nil
}

// genPacket is a parsed packet ready for code generation.
type genPacket struct {
	Title    string
	TypeName string
	ID       int
	Fields   []genField
}

// rawProp mirrors a Mojang JSON Schema property (or array element / definition prop).
type rawProp struct {
	Title   string            `json:"title"`
	Type    string            `json:"type"`
	Under   string            `json:"x-underlying-type"`
	Options []string          `json:"x-serialization-options"`
	Ordinal int               `json:"x-ordinal-index"`
	Ref     string            `json:"$ref"`
	Enum    []string          `json:"enum"`
	Items   *rawProp          `json:"items"`
	OneOf   []json.RawMessage `json:"oneOf"`
}

type rawDef struct {
	Title      string             `json:"title"`
	Properties map[string]rawProp `json:"properties"`
}

type rawGenDoc struct {
	Title       string                     `json:"title"`
	Definitions map[string]rawDef          `json:"definitions"`
	Properties  map[string]rawProp         `json:"properties"`
	Meta        map[string]json.RawMessage `json:"$metaProperties"`
}

// parseGenPacket parses a Mojang protocol-docs JSON file into a genPacket.
func parseGenPacket(b []byte) (genPacket, error) {
	var doc rawGenDoc
	if err := json.Unmarshal(b, &doc); err != nil {
		return genPacket{}, fmt.Errorf("unmarshal: %w", err)
	}
	pk := genPacket{Title: doc.Title, TypeName: goTypeName(doc.Title)}
	if raw, ok := doc.Meta["[cereal:packet]"]; ok {
		_ = json.Unmarshal(raw, &pk.ID)
	}

	required := map[string]bool{}
	// "required" is parsed separately because it lives at the top level.
	var top struct {
		Required []string `json:"required"`
	}
	_ = json.Unmarshal(b, &top)
	for _, r := range top.Required {
		required[r] = true
	}

	for name, rp := range doc.Properties {
		f := buildField(name, rp, doc.Definitions)
		f.Required = required[name]
		pk.Fields = append(pk.Fields, f)
	}
	sort.SliceStable(pk.Fields, func(i, j int) bool {
		return pk.Fields[i].Ordinal < pk.Fields[j].Ordinal
	})
	return pk, nil
}

// buildField turns a raw property into a genField, resolving refs and arrays.
func buildField(name string, rp rawProp, defs map[string]rawDef) genField {
	f := genField{
		DocName:  name,
		GoName:   goFieldName(name),
		Ordinal:  rp.Ordinal,
		Under:    rp.Under,
		Options:  rp.Options,
		JSONType: rp.Type,
		Enum:     len(rp.Enum) > 0,
	}
	switch {
	case len(rp.OneOf) > 0:
		f.OneOf = true
	case rp.Type == "array" && rp.Items != nil:
		f.IsArray = true
		elem := buildField(name+"Elem", *rp.Items, defs)
		f.Elem = &elem
	case rp.Ref != "":
		applyRef(&f, rp.Ref, defs)
	}
	return f
}

// applyRef resolves a "#/definitions/<id>" reference into the field, flattening a
// single-scalar wrapper or marking a multi-field definition as composite.
func applyRef(f *genField, ref string, defs map[string]rawDef) {
	key := ref[lastSlash(ref)+1:]
	def, ok := defs[key]
	if !ok {
		return
	}
	f.Ref = ref
	f.RefTitle = def.Title
	if len(def.Properties) == 1 {
		for _, inner := range def.Properties {
			f.Under = inner.Under
			f.Options = inner.Options
			f.JSONType = inner.Type
			f.Enum = len(inner.Enum) > 0
		}
		return
	}
	f.Composite = true
}

func lastSlash(s string) int {
	idx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			idx = i
		}
	}
	return idx
}
