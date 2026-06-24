package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"
)

// DocField is a single packet field as described by Mojang's protocol docs.
type DocField struct {
	// Name is the human field name (the JSON property key), e.g. "Target Actor Runtime ID".
	Name string
	// Ordinal is the field's position on the wire (x-ordinal-index).
	Ordinal int
	// UnderlyingType is the concrete scalar wire type (x-underlying-type), e.g. "uint64", "float".
	// For a $ref to a single-scalar wrapper this is flattened to that scalar's type.
	UnderlyingType string
	// Options are the x-serialization-options, e.g. ["Compression"], ["Enum-as-Value"].
	Options []string
	// JSONType is the JSON Schema "type" (string/integer/number/boolean/object).
	JSONType string
	// RefTitle is the title of a referenced definition, if this field was a $ref.
	RefTitle string
	// Composite is true when this field references a multi-field definition (e.g. BlockPos, Vec3).
	Composite bool
	// Enum is true when the field declares an enum. An enum without the
	// "Enum-as-Value" option is serialized as its string name on the wire.
	Enum bool
}

// HasOption reports whether the field carries the given x-serialization-option.
func (f DocField) HasOption(o string) bool {
	return slices.Contains(f.Options, o)
}

// DocPacket is a packet as described by Mojang's protocol docs.
type DocPacket struct {
	// Name is the packet title, e.g. "AnimatePacket".
	Name string
	// ID is the numeric packet id from $metaProperties["[cereal:packet]"].
	ID int
	// Fields are the packet's fields, sorted by Ordinal.
	Fields []DocField
}

// rawField mirrors a Mojang JSON Schema property.
type rawField struct {
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Under   string   `json:"x-underlying-type"`
	Options []string `json:"x-serialization-options"`
	Ordinal int      `json:"x-ordinal-index"`
	Ref     string   `json:"$ref"`
	Enum    []string `json:"enum"`
}

// rawType mirrors a Mojang JSON Schema definition (a reusable composite type).
type rawType struct {
	Title      string              `json:"title"`
	Properties map[string]rawField `json:"properties"`
}

// rawDoc mirrors the top level of a Mojang protocol-docs JSON file.
type rawDoc struct {
	Title       string                     `json:"title"`
	Definitions map[string]rawType         `json:"definitions"`
	Properties  map[string]rawField        `json:"properties"`
	Meta        map[string]json.RawMessage `json:"$metaProperties"`
}

// parseDocPacket parses a single Mojang protocol-docs JSON file into a DocPacket.
func parseDocPacket(b []byte) (DocPacket, error) {
	var doc rawDoc
	if err := json.Unmarshal(b, &doc); err != nil {
		return DocPacket{}, fmt.Errorf("unmarshal doc: %w", err)
	}
	pk := DocPacket{Name: doc.Title}

	if raw, ok := doc.Meta["[cereal:packet]"]; ok {
		if err := json.Unmarshal(raw, &pk.ID); err != nil {
			return DocPacket{}, fmt.Errorf("parse [cereal:packet] id: %w", err)
		}
	}

	for name, rf := range doc.Properties {
		field := DocField{
			Name:           name,
			Ordinal:        rf.Ordinal,
			UnderlyingType: rf.Under,
			Options:        rf.Options,
			JSONType:       rf.Type,
			Enum:           len(rf.Enum) > 0,
		}
		if rf.Ref != "" {
			resolveRef(&field, rf.Ref, doc.Definitions)
		}
		pk.Fields = append(pk.Fields, field)
	}
	sort.SliceStable(pk.Fields, func(i, j int) bool {
		return pk.Fields[i].Ordinal < pk.Fields[j].Ordinal
	})
	return pk, nil
}

// resolveRef resolves a "#/definitions/<id>" reference into the field. A wrapper
// definition with a single scalar property is flattened to that scalar (Mojang
// wraps things like runtime ids in single-field objects, but gophertunnel writes
// them as one scalar op). A multi-property definition is marked Composite.
func resolveRef(field *DocField, ref string, defs map[string]rawType) {
	key := ref[strings.LastIndexByte(ref, '/')+1:]
	def, ok := defs[key]
	if !ok {
		return
	}
	field.RefTitle = def.Title
	if len(def.Properties) == 1 {
		for _, inner := range def.Properties {
			field.UnderlyingType = inner.Under
			field.Options = inner.Options
			field.JSONType = inner.Type
		}
		return
	}
	field.Composite = true
}
