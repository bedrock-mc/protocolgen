// Package updateguide turns a human-readable protocol changelog and the target
// Mojang schema snapshot into implementation-oriented gophertunnel snippets.
package updateguide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"protocolgen/internal/naming"
)

// Generate reads changed schema names from changelog and renders their target
// definitions from schemasDir. The changelog remains the source of truth for
// which definitions changed; unchanged schemas are never emitted.
func Generate(changelog []byte, schemasDir string) ([]byte, error) {
	documents, err := loadDocuments(schemasDir)
	if err != nil {
		return nil, err
	}
	parsed, err := parseChangelog(string(changelog))
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	fmt.Fprintf(&output, "# gophertunnel update guide — protocol %s to %s\n\n", parsed.fromProtocol, parsed.toProtocol)
	if parsed.versions != "" {
		fmt.Fprintf(&output, "%s\n\n", parsed.versions)
	}
	if len(parsed.table) != 0 {
		for _, line := range parsed.table {
			fmt.Fprintln(&output, line)
		}
		fmt.Fprintln(&output)
	}
	fmt.Fprintln(&output, "Read `changelog.md` for the human-readable diff. This guide gives a target-version `Marshal` transcription for each changed definition.")
	fmt.Fprintln(&output)
	fmt.Fprintln(&output, "**The Go below is a transcription aid, not a patch.** Names come from the schema, so they may not match gophertunnel's existing names. Field comments are emitted only when Mojang provides a description.")

	for _, section := range parsed.sections {
		fmt.Fprintf(&output, "\n## %s\n", section.title)
		for _, item := range section.items {
			fmt.Fprintf(&output, "\n### %s\n", item.name)
			if item.summary != "" {
				fmt.Fprintf(&output, "%s\n", item.summary)
			}
			if len(item.changes) != 0 {
				fmt.Fprintln(&output)
				for _, change := range item.changes {
					fmt.Fprintln(&output, change)
				}
			}
			if len(item.affects) != 0 {
				fmt.Fprintf(&output, "\n**Also update:** %s — %s embeds this type.\n", quotedNames(item.affects), cardinality(len(item.affects)))
			}
			if section.removed {
				fmt.Fprintln(&output, "\nRemove the corresponding target definition and its registrations or references.")
				continue
			}
			document, ok := documents.byTitle[item.name]
			if !ok {
				return nil, fmt.Errorf("changed schema %q is not present in %s", item.name, schemasDir)
			}
			snippet, err := renderSnippet(document, documents, section.kind)
			if err != nil {
				return nil, fmt.Errorf("render %s: %w", item.name, err)
			}
			snippet, err = formatSnippet(snippet)
			if err != nil {
				return nil, fmt.Errorf("format %s: %w", item.name, err)
			}
			fmt.Fprintf(&output, "\n```go\n%s\n```\n", snippet)
		}
	}
	return output.Bytes(), nil
}

func formatSnippet(snippet string) (string, error) {
	formatted, err := format.Source([]byte("package guide\n\n" + snippet + "\n"))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.TrimPrefix(string(formatted), "package guide\n\n")), nil
}

type schemaSet struct {
	byFile  map[string]map[string]any
	byTitle map[string]map[string]any
}

func loadDocuments(directory string) (schemaSet, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return schemaSet{}, fmt.Errorf("read target schema directory: %w", err)
	}
	set := schemaSet{byFile: map[string]map[string]any{}, byTitle: map[string]map[string]any{}}
	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), "__") || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(directory, entry.Name()))
		if err != nil {
			return schemaSet{}, fmt.Errorf("read target schema %s: %w", entry.Name(), err)
		}
		var document map[string]any
		if err := json.Unmarshal(data, &document); err != nil {
			return schemaSet{}, fmt.Errorf("parse target schema %s: %w", entry.Name(), err)
		}
		set.byFile[entry.Name()] = document
		set.byTitle[strings.TrimSuffix(entry.Name(), ".json")] = document
		if title, _ := document["title"].(string); title != "" {
			set.byTitle[title] = document
		}
	}
	return set, nil
}

type changelogDocument struct {
	fromProtocol string
	toProtocol   string
	versions     string
	table        []string
	sections     []changelogSection
}

type changelogSection struct {
	title   string
	kind    string
	removed bool
	items   []changelogItem
}

type changelogItem struct {
	name    string
	summary string
	changes []string
	affects []string
}

func parseChangelog(input string) (changelogDocument, error) {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	document := changelogDocument{}
	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			break
		}
		if strings.HasPrefix(line, "# ") {
			parts := strings.Fields(line)
			for index, part := range parts {
				if part == "to" && index > 0 && index+1 < len(parts) {
					document.fromProtocol = strings.Trim(parts[index-1], "` ")
					document.toProtocol = strings.Trim(parts[index+1], "` ")
				}
			}
			continue
		}
		if document.versions == "" && strings.Contains(line, "→") && !strings.HasPrefix(line, "|") {
			document.versions = strings.TrimSpace(line)
		}
		if strings.HasPrefix(line, "|") {
			document.table = append(document.table, line)
		}
	}
	if document.fromProtocol == "" || document.toProtocol == "" {
		return changelogDocument{}, fmt.Errorf("changelog heading must identify the source and target protocols")
	}

	var section *changelogSection
	var item *changelogItem
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		switch {
		case strings.HasPrefix(line, "## "):
			title := strings.TrimPrefix(line, "## ")
			kind := ""
			switch {
			case strings.Contains(title, "Packets"):
				kind = "packet"
			case strings.Contains(title, "Types"):
				kind = "type"
			case strings.Contains(title, "Enums"):
				kind = "enum"
			}
			if kind == "" {
				section, item = nil, nil
				continue
			}
			document.sections = append(document.sections, changelogSection{title: title, kind: kind, removed: strings.Contains(title, "Removed")})
			section = &document.sections[len(document.sections)-1]
			item = nil
		case strings.HasPrefix(line, "### ") && section != nil:
			section.items = append(section.items, changelogItem{name: strings.Trim(strings.TrimPrefix(line, "### "), "`")})
			item = &section.items[len(section.items)-1]
		case item != nil && strings.HasPrefix(line, "**Also affects:**"):
			item.affects = splitNames(strings.TrimSpace(strings.TrimPrefix(line, "**Also affects:**")))
		case item != nil && strings.HasPrefix(line, "-"):
			item.changes = append(item.changes, line)
		case item != nil && line != "" && item.summary == "" && !strings.HasPrefix(line, "|"):
			item.summary = line
		}
	}
	return document, nil
}

func splitNames(input string) []string {
	parts := strings.Split(input, ",")
	names := make([]string, 0, len(parts))
	for _, part := range parts {
		if name := strings.Trim(strings.TrimSpace(part), "`"); name != "" {
			names = append(names, name)
		}
	}
	return names
}

func quotedNames(names []string) string {
	quoted := make([]string, len(names))
	for index, name := range names {
		quoted[index] = "`" + name + "`"
	}
	return strings.Join(quoted, ", ")
}

func cardinality(count int) string {
	if count == 1 {
		return "it"
	}
	return "each"
}

func renderSnippet(document map[string]any, schemas schemaSet, kind string) (string, error) {
	if kind == "enum" {
		return renderEnum(document)
	}
	name := schemaName(document)
	body, err := resolveRoot(document, schemas, map[string]bool{})
	if err != nil {
		return "", err
	}
	fields, err := schemaFields(body, schemas)
	if err != nil {
		return "", err
	}
	var output strings.Builder
	fmt.Fprintf(&output, "type %s struct {\n", name)
	for _, field := range fields {
		if field.description != "" {
			fmt.Fprintf(&output, "\t// %s %s\n", field.name, oneLine(field.description))
		}
		fmt.Fprintf(&output, "\t%s %s\n", field.name, field.goType)
	}
	fmt.Fprintln(&output, "}")
	if kind == "packet" {
		fmt.Fprintf(&output, "\nfunc (*%s) ID() uint32 {\n\treturn ID%s\n}\n", name, name)
	}
	fmt.Fprintf(&output, "\nfunc (pk *%s) Marshal(io protocol.IO) {\n", name)
	for _, field := range fields {
		fmt.Fprintf(&output, "\t%s\n", field.marshal)
	}
	fmt.Fprint(&output, "}")
	return output.String(), nil
}

func renderEnum(document map[string]any) (string, error) {
	values, ok := document["enum"].([]any)
	if !ok || len(values) == 0 {
		return "", fmt.Errorf("enum has no values")
	}
	explicit, _ := document["x-enum-values"].([]any)
	name := schemaName(document)
	var output strings.Builder
	fmt.Fprintln(&output, "const (")
	for index, rawValue := range values {
		valueName := naming.GoExportName(fmt.Sprint(rawValue))
		value := int64(index)
		if index < len(explicit) {
			parsed, err := number(explicit[index])
			if err != nil {
				return "", fmt.Errorf("enum value %s: %w", valueName, err)
			}
			value = parsed
		}
		fmt.Fprintf(&output, "\t%s%s = %d\n", name, valueName, value)
	}
	fmt.Fprint(&output, ")")
	return output.String(), nil
}

type renderedField struct {
	ordinal     int
	name        string
	description string
	goType      string
	marshal     string
}

func schemaFields(document map[string]any, schemas schemaSet) ([]renderedField, error) {
	properties, ok := document["properties"].(map[string]any)
	if !ok {
		if document["type"] == "object" {
			return nil, nil
		}
		return nil, fmt.Errorf("schema is not an object")
	}
	required := map[string]bool{}
	if values, ok := document["required"].([]any); ok {
		for _, value := range values {
			required[fmt.Sprint(value)] = true
		}
	}
	fields := make([]renderedField, 0, len(properties))
	for rawName, rawSchema := range properties {
		schema, ok := rawSchema.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("field %s is not a schema", rawName)
		}
		ordinal, err := number(schema["x-ordinal-index"])
		if err != nil {
			return nil, fmt.Errorf("field %s has no ordinal", rawName)
		}
		name := naming.GoExportName(rawName)
		goType, marshal, marshaler, err := renderFieldSchema(schema, schemas, "&pk."+name)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", rawName, err)
		}
		if !required[rawName] {
			valueType := goType
			goType = "protocol.Optional[" + valueType + "]"
			if marshaler {
				marshal = "protocol.OptionalMarshaler(io, &pk." + name + ")"
			} else {
				address := "&pk." + name
				funcSlicePrefix := "protocol.FuncSlice(io, " + address + ", "
				sliceCall := "protocol.Slice(io, " + address + ")"
				switch {
				case strings.HasPrefix(marshal, funcSlicePrefix) && strings.HasSuffix(marshal, ")"):
					codec := strings.TrimSuffix(strings.TrimPrefix(marshal, funcSlicePrefix), ")")
					marshal = "protocol.OptionalFunc(io, " + address + ", func(value *" + valueType + ") { protocol.FuncSlice(io, value, " + codec + ") })"
				case marshal == sliceCall:
					marshal = "protocol.OptionalFunc(io, " + address + ", func(value *" + valueType + ") { protocol.Slice(io, value) })"
				default:
					marshal = strings.TrimSuffix(marshal, "("+address+")")
					marshal = "protocol.OptionalFunc(io, " + address + ", " + marshal + ")"
				}
			}
		}
		description, _ := schema["description"].(string)
		fields = append(fields, renderedField{ordinal: int(ordinal), name: name, description: description, goType: goType, marshal: marshal})
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].ordinal < fields[j].ordinal })
	return fields, nil
}

// renderFieldSchema returns the Go type, either a complete marshal statement
// or an IO method expression for optionals, and whether the type marshals itself.
func renderFieldSchema(schema map[string]any, schemas schemaSet, address string) (string, string, bool, error) {
	if reference, _ := schema["$ref"].(string); reference != "" {
		target, err := resolveReference(reference, schemas)
		if err != nil {
			return "", "", false, err
		}
		if _, enum := target["enum"]; enum {
			_, method, err := primitive(target)
			return schemaName(target), method, false, err
		}
		return schemaName(target), "protocol.Single(io, " + address + ")", true, nil
	}
	switch schema["type"] {
	case "string":
		return "string", chooseStatement(address, "io.String"), false, nil
	case "boolean":
		return "bool", chooseStatement(address, "io.Bool"), false, nil
	case "integer", "number":
		goType, method, err := primitive(schema)
		return goType, chooseStatement(address, method), false, err
	case "array":
		items, ok := schema["items"].(map[string]any)
		if !ok {
			return "", "", false, fmt.Errorf("array has no item schema")
		}
		itemType, itemMarshal, itemMarshaler, err := renderFieldSchema(items, schemas, "")
		if err != nil {
			return "", "", false, err
		}
		if itemMarshaler {
			return "[]" + itemType, "protocol.Slice(io, " + address + ")", false, nil
		}
		return "[]" + itemType, "protocol.FuncSlice(io, " + address + ", " + itemMarshal + ")", false, nil
	case "object":
		return schemaName(schema), "protocol.Single(io, " + address + ")", true, nil
	default:
		return "", "", false, fmt.Errorf("unsupported schema type %q", schema["type"])
	}
}

func chooseStatement(address, method string) string {
	if address == "" {
		return method
	}
	return method + "(" + address + ")"
}

func primitive(schema map[string]any) (string, string, error) {
	underlying := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(schema["x-underlying-type"]), "_", ""), " ", ""))
	options := strings.ToLower(fmt.Sprint(schema["x-serialization-options"]))
	compressed := strings.Contains(options, "compression")
	bigEndian := strings.Contains(options, "big endian")
	if underlying == "" || underlying == "<nil>" {
		underlying = fmt.Sprint(schema["type"])
	}
	if compressed {
		switch underlying {
		case "int", "int32", "short", "int16", "byte", "int8":
			return "int32", "io.Varint32", nil
		case "uint", "uint32", "ushort", "uint16", "ubyte", "uint8":
			return "uint32", "io.Varuint32", nil
		case "long", "int64":
			return "int64", "io.Varint64", nil
		case "ulong", "uint64":
			return "uint64", "io.Varuint64", nil
		}
	}
	mapping := map[string][2]string{
		"byte": {"int8", "io.Int8"}, "int8": {"int8", "io.Int8"},
		"ubyte": {"uint8", "io.Uint8"}, "uint8": {"uint8", "io.Uint8"},
		"short": {"int16", "io.Int16"}, "int16": {"int16", "io.Int16"},
		"ushort": {"uint16", "io.Uint16"}, "uint16": {"uint16", "io.Uint16"},
		"int": {"int32", "io.Int32"}, "int32": {"int32", "io.Int32"},
		"uint": {"uint32", "io.Uint32"}, "uint32": {"uint32", "io.Uint32"},
		"long": {"int64", "io.Int64"}, "int64": {"int64", "io.Int64"},
		"ulong": {"uint64", "io.Uint64"}, "uint64": {"uint64", "io.Uint64"},
		"float": {"float32", "io.Float32"}, "float32": {"float32", "io.Float32"},
		"double": {"float64", "io.Float64"}, "float64": {"float64", "io.Float64"},
	}
	result, ok := mapping[underlying]
	if !ok {
		return "", "", fmt.Errorf("unsupported primitive %q", underlying)
	}
	if bigEndian {
		if result[1] != "io.Int32" {
			return "", "", fmt.Errorf("unsupported big-endian primitive %q", underlying)
		}
		result[1] = "io.BEInt32"
	}
	return result[0], result[1], nil
}

func resolveRoot(document map[string]any, schemas schemaSet, active map[string]bool) (map[string]any, error) {
	reference, _ := document["$ref"].(string)
	if reference == "" {
		return document, nil
	}
	if active[reference] {
		return nil, fmt.Errorf("cyclic root reference %s", reference)
	}
	active[reference] = true
	target, err := resolveReference(reference, schemas)
	if err != nil {
		return nil, err
	}
	return resolveRoot(target, schemas, active)
}

func resolveReference(reference string, schemas schemaSet) (map[string]any, error) {
	file := filepath.Base(strings.SplitN(reference, "#", 2)[0])
	target, ok := schemas.byFile[file]
	if !ok {
		return nil, fmt.Errorf("missing reference %s", reference)
	}
	return target, nil
}

func schemaName(document map[string]any) string {
	title, _ := document["title"].(string)
	return naming.GoExportName(naming.PublicTypeName(title))
}

func number(value any) (int64, error) {
	switch value := value.(type) {
	case float64:
		return int64(value), nil
	case int:
		return int64(value), nil
	case int64:
		return value, nil
	case json.Number:
		return value.Int64()
	case string:
		return strconv.ParseInt(value, 10, 64)
	default:
		return 0, fmt.Errorf("not an integer")
	}
}

func oneLine(value string) string {
	return strings.Join(strings.Fields(value), " ")
}
