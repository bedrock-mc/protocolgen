// Package nbtencoding applies the reviewed NBT wire-format overlay.
package nbtencoding

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"protocolgen/internal/manifest"
)

const SchemaVersion uint32 = 1

type Table struct {
	SchemaVersion uint32          `json:"schema_version"`
	Target        manifest.Target `json:"target"`
	Source        Source          `json:"source"`
	Fields        []Entry         `json:"fields"`
}

type Source struct {
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Locator    string `json:"locator"`
	SHA256     string `json:"sha256"`
}

type Entry struct {
	PacketID     uint32               `json:"packet_id"`
	PacketName   string               `json:"packet_name"`
	FieldOrdinal int                  `json:"field_ordinal"`
	FieldName    string               `json:"field_name"`
	Path         string               `json:"path"`
	Encoding     manifest.NBTEncoding `json:"encoding"`
	Evidence     Evidence             `json:"evidence"`
}

type Evidence struct {
	Locator string `json:"locator"`
	Summary string `json:"summary"`
}

func Load(path string) (Table, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Table{}, fmt.Errorf("read NBT encoding table: %w", err)
	}
	var table Table
	if err := json.Unmarshal(data, &table); err != nil {
		return Table{}, fmt.Errorf("parse NBT encoding table: %w", err)
	}
	if err := table.validateShape(); err != nil {
		return Table{}, fmt.Errorf("validate NBT encoding table: %w", err)
	}
	return table, nil
}

func (t Table) Apply(value *manifest.Manifest) error {
	if value == nil {
		return fmt.Errorf("cannot apply NBT encoding table to nil manifest")
	}
	if err := t.Validate(value.Target, value.Packets); err != nil {
		return err
	}
	for packetIndex := range value.Packets {
		packet := &value.Packets[packetIndex]
		for fieldIndex := range packet.Fields {
			field := &packet.Fields[fieldIndex]
			for _, entry := range t.Fields {
				if entry.PacketID != packet.ID || entry.FieldOrdinal != field.Ordinal {
					continue
				}
				node, err := nodeAt(&field.Encode, entry.Path)
				if err != nil {
					return fmt.Errorf("NBT field %d.%d path %q: %w", packet.ID, field.Ordinal, entry.Path, err)
				}
				node.Encoding = string(entry.Encoding)
			}
		}
	}
	return manifest.Validate(*value)
}

func (t Table) Validate(target manifest.Target, packets []manifest.Packet) error {
	if err := t.validateShape(); err != nil {
		return err
	}
	if t.Target != target {
		return fmt.Errorf("NBT encoding table target does not match manifest target")
	}
	packetByID := make(map[uint32]*manifest.Packet, len(packets))
	for index := range packets {
		packetByID[packets[index].ID] = &packets[index]
	}
	seen := make(map[string]bool, len(t.Fields))
	for _, entry := range t.Fields {
		packet, ok := packetByID[entry.PacketID]
		if !ok {
			return fmt.Errorf("NBT encoding table references unknown packet id %d", entry.PacketID)
		}
		if packet.Name != entry.PacketName {
			return fmt.Errorf("NBT encoding table packet id %d names %q, manifest names %q", entry.PacketID, entry.PacketName, packet.Name)
		}
		fieldIndex := -1
		for index := range packet.Fields {
			if packet.Fields[index].Ordinal == entry.FieldOrdinal {
				fieldIndex = index
				break
			}
		}
		if fieldIndex < 0 {
			return fmt.Errorf("NBT encoding table references unknown field %d.%d", entry.PacketID, entry.FieldOrdinal)
		}
		field := packet.Fields[fieldIndex]
		if field.Name != entry.FieldName {
			return fmt.Errorf("NBT encoding table field %d.%d names %q, manifest names %q", entry.PacketID, entry.FieldOrdinal, entry.FieldName, field.Name)
		}
		node, err := nodeAt(&field.Encode, entry.Path)
		if err != nil {
			return fmt.Errorf("NBT encoding table field %d.%d path %q: %w", entry.PacketID, entry.FieldOrdinal, entry.Path, err)
		}
		if !manifest.IsNBT(*node) {
			return fmt.Errorf("NBT encoding table field %d.%d path %q does not name an NBT node", entry.PacketID, entry.FieldOrdinal, entry.Path)
		}
		key := fmt.Sprintf("%d/%d/%s", entry.PacketID, entry.FieldOrdinal, entry.Path)
		if seen[key] {
			return fmt.Errorf("duplicate NBT encoding entry %s", key)
		}
		seen[key] = true
	}

	for _, packet := range packets {
		for _, field := range packet.Fields {
			var nbtPaths []string
			walkNBT(field.Encode, "encode", &nbtPaths)
			for _, path := range nbtPaths {
				key := fmt.Sprintf("%d/%d/%s", packet.ID, field.Ordinal, path)
				if !seen[key] {
					return fmt.Errorf("missing NBT encoding entry for packet %d field %d path %q", packet.ID, field.Ordinal, path)
				}
			}
		}
	}
	return nil
}

func (t Table) validateShape() error {
	if t.SchemaVersion != SchemaVersion {
		return fmt.Errorf("schema_version %d is not v%d", t.SchemaVersion, SchemaVersion)
	}
	if t.Target.MinecraftVersion == "" || t.Target.ProtocolVersion <= 0 {
		return fmt.Errorf("target must pin minecraft_version and positive protocol_version")
	}
	if t.Source.Repository == "" || t.Source.Revision == "" || t.Source.Locator == "" || t.Source.SHA256 == "" {
		return fmt.Errorf("source is missing repository, revision, locator, or sha256")
	}
	if len(t.Source.Revision) != 40 || !isHex(t.Source.Revision) {
		return fmt.Errorf("source revision is not a full SHA-1")
	}
	if len(t.Source.SHA256) != len("sha256:")+64 || !strings.HasPrefix(t.Source.SHA256, "sha256:") || !isHex(strings.TrimPrefix(t.Source.SHA256, "sha256:")) {
		return fmt.Errorf("source sha256 is not a SHA-256 fingerprint")
	}
	for index, entry := range t.Fields {
		if entry.PacketID == 0 || entry.PacketName == "" || entry.FieldOrdinal < 0 || entry.FieldName == "" || entry.Path == "" {
			return fmt.Errorf("fields[%d] is incomplete", index)
		}
		if !manifest.ValidNBTEncoding(string(entry.Encoding)) {
			return fmt.Errorf("fields[%d] has invalid NBT encoding %q", index, entry.Encoding)
		}
		if entry.Evidence.Locator == "" || entry.Evidence.Summary == "" {
			return fmt.Errorf("fields[%d] has incomplete evidence", index)
		}
	}
	return nil
}

func nodeAt(root *manifest.Node, path string) (*manifest.Node, error) {
	if path == "encode" {
		return root, nil
	}
	if !strings.HasPrefix(path, "encode.") {
		return nil, fmt.Errorf("path must start with encode")
	}
	path = strings.TrimPrefix(path, "encode.")
	node := root
	for path != "" {
		part, rest := nextPathPart(path)
		switch {
		case part == "prefix", part == "element", part == "value", part == "key", part == "control", part == "default":
			child := childNode(*node, part)
			if child == nil {
				return nil, fmt.Errorf("node has no %s child", part)
			}
			node = child
		case strings.HasPrefix(part, "fields["):
			index, err := indexedPart(part, "fields")
			if err != nil || index < 0 || index >= len(node.Fields) {
				return nil, fmt.Errorf("invalid field selector %q", part)
			}
			if !strings.HasPrefix(rest, ".encode") {
				return nil, fmt.Errorf("field selector %q must select encode", part)
			}
			node = &node.Fields[index].Encode
			rest = strings.TrimPrefix(rest, ".encode")
		case strings.HasPrefix(part, "elements["):
			index, err := indexedPart(part, "elements")
			if err != nil || index < 0 || index >= len(node.Elements) {
				return nil, fmt.Errorf("invalid element selector %q", part)
			}
			node = &node.Elements[index]
		case strings.HasPrefix(part, "variants["):
			index, err := indexedPart(part, "variants")
			if err != nil || index < 0 || index >= len(node.Variants) {
				return nil, fmt.Errorf("invalid variant selector %q", part)
			}
			if !strings.HasPrefix(rest, ".encode") {
				return nil, fmt.Errorf("variant selector %q must select encode", part)
			}
			node = &node.Variants[index].Encode
			rest = strings.TrimPrefix(rest, ".encode")
		case strings.HasPrefix(part, "cases["):
			caseIndex, err := indexedPart(part, "cases")
			if err != nil || caseIndex < 0 || caseIndex >= len(node.Cases) {
				return nil, fmt.Errorf("invalid case selector %q", part)
			}
			if !strings.HasPrefix(rest, ".encode[") {
				return nil, fmt.Errorf("case selector %q must select encode", part)
			}
			rest = strings.TrimPrefix(rest, ".encode[")
			close := strings.IndexByte(rest, ']')
			if close < 0 {
				return nil, fmt.Errorf("invalid case encode selector in %q", part)
			}
			nodeIndex, err := strconv.Atoi(rest[:close])
			if err != nil || nodeIndex < 0 || nodeIndex >= len(node.Cases[caseIndex].Encode) {
				return nil, fmt.Errorf("invalid case encode selector in %q", part)
			}
			node = &node.Cases[caseIndex].Encode[nodeIndex]
			rest = rest[close+1:]
		default:
			return nil, fmt.Errorf("unknown path component %q", part)
		}
		path = strings.TrimPrefix(rest, ".")
	}
	return node, nil
}

func nextPathPart(path string) (string, string) {
	if dot := strings.IndexByte(path, '.'); dot >= 0 {
		return path[:dot], path[dot:]
	}
	return path, ""
}

func indexedPart(part, name string) (int, error) {
	prefix := name + "["
	if !strings.HasPrefix(part, prefix) || !strings.HasSuffix(part, "]") {
		return 0, fmt.Errorf("invalid selector")
	}
	return strconv.Atoi(part[len(prefix) : len(part)-1])
}

func childNode(node manifest.Node, name string) *manifest.Node {
	switch name {
	case "prefix":
		return node.Prefix
	case "element":
		return node.Element
	case "value":
		return node.Value
	case "key":
		return node.Key
	case "control":
		return node.Control
	case "default":
		return node.Default
	default:
		return nil
	}
}

func walkNBT(node manifest.Node, path string, result *[]string) {
	if manifest.IsNBT(node) {
		*result = append(*result, path)
	}
	if node.Prefix != nil {
		walkNBT(*node.Prefix, path+".prefix", result)
	}
	if node.Element != nil {
		walkNBT(*node.Element, path+".element", result)
	}
	if node.Value != nil {
		walkNBT(*node.Value, path+".value", result)
	}
	if node.Key != nil {
		walkNBT(*node.Key, path+".key", result)
	}
	if node.Control != nil {
		walkNBT(*node.Control, path+".control", result)
	}
	if node.Default != nil {
		walkNBT(*node.Default, path+".default", result)
	}
	for index := range node.Elements {
		walkNBT(node.Elements[index], fmt.Sprintf("%s.elements[%d]", path, index), result)
	}
	for index := range node.Fields {
		walkNBT(node.Fields[index].Encode, fmt.Sprintf("%s.fields[%d].encode", path, index), result)
	}
	for index := range node.Variants {
		walkNBT(node.Variants[index].Encode, fmt.Sprintf("%s.variants[%d].encode", path, index), result)
	}
	for caseIndex := range node.Cases {
		for nodeIndex := range node.Cases[caseIndex].Encode {
			walkNBT(node.Cases[caseIndex].Encode[nodeIndex], fmt.Sprintf("%s.cases[%d].encode[%d]", path, caseIndex, nodeIndex), result)
		}
	}
}

func isHex(value string) bool {
	_, err := hex.DecodeString(value)
	return err == nil
}
