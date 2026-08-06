package ingest

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"protocolgen/internal/claims"
	"protocolgen/internal/manifest"
)

type endstoneLowerer struct {
	types  map[string]any
	enums  map[string]any
	active map[string]bool
	source manifest.SourcePin
}

func ParseEndstone(root string, pin manifest.SourcePin, corrections string) (claims.Result, error) {
	packetDocs, err := loadJSONFiles(filepath.Join(root, "packets"))
	if err != nil {
		return claims.Result{}, fmt.Errorf("load Endstone packets: %w", err)
	}
	typeDocs, err := loadJSONFiles(filepath.Join(root, "types"))
	if err != nil {
		return claims.Result{}, fmt.Errorf("load Endstone types: %w", err)
	}
	enumDocs, err := loadJSONFiles(filepath.Join(root, "enums"))
	if err != nil {
		return claims.Result{}, fmt.Errorf("load Endstone enums: %w", err)
	}
	combined := map[string]any{}
	for name, value := range packetDocs {
		combined["packets/"+name] = value
	}
	for name, value := range typeDocs {
		combined["types/"+name] = value
	}
	for name, value := range enumDocs {
		combined["enums/"+name] = value
	}
	proofs, err := applyCorrections(combined, corrections, pin)
	if err != nil {
		return claims.Result{}, err
	}
	for key, value := range combined {
		switch {
		case strings.HasPrefix(key, "packets/"):
			packetDocs[strings.TrimPrefix(key, "packets/")] = value
		case strings.HasPrefix(key, "types/"):
			typeDocs[strings.TrimPrefix(key, "types/")] = value
		case strings.HasPrefix(key, "enums/"):
			enumDocs[strings.TrimPrefix(key, "enums/")] = value
		}
	}
	target, err := endstoneTarget(root, pin)
	if err != nil {
		return claims.Result{}, err
	}
	lowerer := &endstoneLowerer{types: typeDocs, enums: enumDocs, active: map[string]bool{}, source: pin}
	type packetDoc struct {
		name string
		id   uint32
	}
	var packets []packetDoc
	for file, rawDocument := range packetDocs {
		document, ok := asMap(rawDocument)
		if !ok {
			return claims.Result{}, fmt.Errorf("Endstone packet %s is not an object", file)
		}
		idValue, ok := asInt(document["id"])
		if !ok || idValue < 0 {
			return claims.Result{}, fmt.Errorf("Endstone packet %s has invalid id", file)
		}
		name := asString(document["name"])
		if name == "" {
			name = strings.TrimSuffix(file, filepath.Ext(file))
		}
		packets = append(packets, packetDoc{name: name, id: uint32(idValue)})
	}
	sort.Slice(packets, func(i, j int) bool {
		if packets[i].id != packets[j].id {
			return packets[i].id < packets[j].id
		}
		return packets[i].name < packets[j].name
	})
	if len(packets) == 0 {
		return claims.Result{}, fmt.Errorf("Endstone source contains no packets")
	}
	result := claims.Result{Pin: pin, Target: target, Overrides: proofs}
	for _, packet := range packets {
		document := packetDocs[packet.name+".json"]
		if document == nil {
			for file, candidate := range packetDocs {
				object, _ := asMap(candidate)
				if asString(object["name"]) == packet.name {
					document = candidate
					_ = file
					break
				}
			}
		}
		object, ok := asMap(document)
		if !ok {
			return claims.Result{}, fmt.Errorf("Endstone packet %s document disappeared", packet.name)
		}
		fields, ok := asArray(object["fields"])
		if !ok {
			return claims.Result{}, fmt.Errorf("Endstone packet %s has no fields array", packet.name)
		}
		for index, rawField := range fields {
			field, ok := asMap(rawField)
			if !ok {
				return claims.Result{}, fmt.Errorf("Endstone packet %s field %d is not an object", packet.name, index)
			}
			name := asString(field["name"])
			if name == "" {
				name = fmt.Sprintf("constant_%d", index)
			}
			typeValue, ok := field["type"]
			if !ok {
				return claims.Result{}, fmt.Errorf("Endstone packet %s field %s has no type", packet.name, name)
			}
			node := lowerer.lowerTypeValue(typeValue, packet.name+safeIdentifierName(name), packet.name+"."+name)
			if field["optional"] == true {
				node = manifest.Optional(node)
			}
			if field["double_optional"] == true {
				node = manifest.Optional(manifest.Optional(node))
			}
			reserved, ignored, compatibility := fieldCompatibility(field)
			if reserved {
				node = manifest.Reserved(node)
			}
			if ignored {
				node = manifest.Ignored(node)
			}
			claim := makeClaim(pin, packet.id, packet.name, parseDirection(object), index, name, field, node, "packets/"+packet.name+".json#/fields/"+fmt.Sprint(index))
			claim.Reserved, claim.Ignored, claim.Compatibility = reserved, ignored, compatibility
			result.Claims = append(result.Claims, claim)
		}
	}
	return result, nil
}

func endstoneTarget(root string, pin manifest.SourcePin) (manifest.Target, error) {
	target := manifest.Target{MinecraftVersion: pin.MinecraftVersion, ProtocolVersion: pin.ProtocolVersion}
	readme, err := os.ReadFile(filepath.Join(root, "README.md"))
	if err == nil {
		for _, line := range strings.Split(string(readme), "\n") {
			if strings.Contains(line, "Minecraft Version:") {
				value := strings.TrimSpace(strings.Trim(strings.SplitN(line, "Minecraft Version:", 2)[1], " *`"))
				target.MinecraftVersion = strings.Join(strings.Split(value, ".")[:minInt(3, len(strings.Split(value, ".")))], ".")
			}
			if strings.Contains(line, "Network Version:") {
				value := strings.TrimSpace(strings.Trim(strings.SplitN(line, "Network Version:", 2)[1], " *`"))
				if parsed, ok := asInt(value); ok {
					target.ProtocolVersion = int(parsed)
				}
			}
		}
	}
	if target.MinecraftVersion != pin.MinecraftVersion || target.ProtocolVersion != pin.ProtocolVersion {
		return manifest.Target{}, fmt.Errorf("Endstone source target %s/%d does not match locked target %s/%d", target.MinecraftVersion, target.ProtocolVersion, pin.MinecraftVersion, pin.ProtocolVersion)
	}
	if target.MinecraftVersion == "" || target.ProtocolVersion == 0 {
		return manifest.Target{}, fmt.Errorf("Endstone source has no complete version metadata")
	}
	return target, nil
}

func (l *endstoneLowerer) lowerNamed(name, context string) manifest.Node {
	if primitiveNode := endstoneScalar(name); primitiveNode.Kind != manifest.KindUnresolved {
		return primitiveNode
	}
	document, ok := namedDocument(l.types, name)
	if !ok {
		return manifest.Unresolved("unknown Endstone type "+name+" at "+context, true)
	}
	if l.active[name] {
		return manifest.Recursive(name)
	}
	object, ok := asMap(document)
	if !ok {
		return manifest.Unresolved("Endstone type "+name+" is not an object", true)
	}
	l.active[name] = true
	result := l.lowerContainer(object, name, context)
	delete(l.active, name)
	result.Semantic = name
	result.TypeID = name
	return result
}

func (l *endstoneLowerer) lowerContainer(object map[string]any, name, context string) manifest.Node {
	values, ok := asArray(object["fields"])
	if !ok {
		return manifest.Unresolved("Endstone type "+name+" has no fields", true)
	}
	fields := make([]manifest.Field, 0, len(values))
	for index, rawField := range values {
		field, ok := asMap(rawField)
		if !ok {
			return manifest.Unresolved(fmt.Sprintf("Endstone type %s field %d is not an object", name, index), true)
		}
		fieldName := asString(field["name"])
		if fieldName == "" {
			fieldName = fmt.Sprintf("constant_%d", index)
		}
		typeValue := field["type"]
		node := l.lowerTypeValue(typeValue, name+safeIdentifierName(fieldName), context+"."+fieldName)
		if field["optional"] == true {
			node = manifest.Optional(node)
		}
		fields = append(fields, manifest.Field{Ordinal: index, Name: fieldName, Semantic: fieldName, Encode: node, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{l.source.ID}}})
	}
	return manifest.Node{Kind: manifest.KindStruct, Semantic: name, TypeID: name, Fields: fields}
}

func (l *endstoneLowerer) lowerTypeValue(value any, hint, context string) manifest.Node {
	if name, ok := value.(string); ok {
		return l.lowerNamed(name, context)
	}
	object, ok := asMap(value)
	if !ok {
		return manifest.Unresolved("Endstone type is not a string or object at "+context, true)
	}
	if switchObject, ok := asMap(object["switch"]); ok {
		controlName := asString(switchObject["type"])
		control := endstoneScalar(controlName)
		enumName := asString(switchObject["enum"])
		discriminants := l.enumValues(enumName)
		cases, ok := asArray(object["cases"])
		if !ok || len(cases) != len(discriminants) {
			return manifest.Unresolved("Endstone switch cases do not have explicit enum ordinals at "+context, true)
		}
		variants := make([]manifest.Variant, 0, len(cases))
		for index, rawCase := range cases {
			value := discriminants[index]
			if rawCase == nil {
				variants = append(variants, manifest.Variant{Value: value, Name: fmt.Sprintf("Empty%d", index), Encode: manifest.Void()})
				continue
			}
			caseName, ok := rawCase.(string)
			if !ok {
				return manifest.Unresolved("Endstone switch case is not a named type at "+context, true)
			}
			variants = append(variants, manifest.Variant{Value: value, Name: caseName, Encode: l.lowerTypeValue(caseName, hint+caseName, context)})
		}
		return manifest.Union(control, variants...)
	}
	if key, hasKey := object["key"]; hasKey {
		valueType, hasValue := object["value"]
		if !hasValue {
			return manifest.Unresolved("Endstone map has no value type at "+context, true)
		}
		return manifest.Map(manifest.Primitive("var_u32"), l.lowerTypeValue(key, hint+"Key", context+".key"), l.lowerTypeValue(valueType, hint+"Value", context+".value"))
	}
	if inner, hasType := object["type"]; hasType {
		var result manifest.Node
		if encoding := asString(object["encoding"]); encoding != "" && strings.EqualFold(asString(inner), "string") {
			result = manifest.String(manifest.Primitive("var_u32"))
		} else {
			result = l.lowerTypeValue(inner, hint, context)
		}
		if enumName := asString(object["enum"]); enumName != "" && result.Kind == manifest.KindPrimitive {
			result = l.lowerEnum(enumName, result, context)
		}
		if repeat, ok := asMap(object["repeat"]); ok {
			result = lowerEndstoneRepeat(repeat, result, context)
		}
		return result
	}
	return manifest.Unresolved("unsupported Endstone type at "+context, true)
}

func (l *endstoneLowerer) lowerEnum(name string, underlying manifest.Node, context string) manifest.Node {
	values := l.enumValuesWithNames(name)
	if len(values) == 0 || underlying.Primitive == nil {
		return manifest.Unresolved("Endstone enum "+name+" is incomplete at "+context, true)
	}
	variants := make([]manifest.Variant, 0, len(values))
	for _, value := range values {
		variants = append(variants, manifest.Variant{Value: value.value, Name: value.name, Encode: manifest.Void()})
	}
	return manifest.Node{Kind: manifest.KindEnum, Primitive: underlying.Primitive, Semantic: name, TypeID: "enums/" + name, Variants: variants}
}

func (l *endstoneLowerer) enumValues(name string) []int64 {
	values := l.enumValuesWithNames(name)
	result := make([]int64, 0, len(values))
	for _, value := range values {
		result = append(result, value.value)
	}
	return result
}

type namedEnumValue struct {
	name  string
	value int64
}

func (l *endstoneLowerer) enumValuesWithNames(name string) []namedEnumValue {
	document, ok := namedDocument(l.enums, name)
	if !ok {
		return nil
	}
	object, ok := asMap(document)
	if !ok {
		return nil
	}
	values, ok := asArray(object["values"])
	if !ok {
		return nil
	}
	result := make([]namedEnumValue, 0, len(values))
	for _, rawValue := range values {
		value, ok := asMap(rawValue)
		if !ok {
			return nil
		}
		number, numberOK := asInt(value["value"])
		nameValue := asString(value["name"])
		if !numberOK || nameValue == "" {
			return nil
		}
		result = append(result, namedEnumValue{name: nameValue, value: number})
	}
	return result
}

func namedDocument(documents map[string]any, name string) (any, bool) {
	if document, ok := documents[name]; ok {
		return document, true
	}
	if document, ok := documents[name+".json"]; ok {
		return document, true
	}
	for file, document := range documents {
		if strings.TrimSuffix(file, filepath.Ext(file)) == name {
			return document, true
		}
	}
	return nil, false
}

func endstoneScalar(name string) manifest.Node {
	switch strings.ToLower(name) {
	case "string", "pstring", "mcpe_string":
		return manifest.String(manifest.Primitive("var_u32"))
	case "restbuffer", "bytearray", "buffer", "bytes":
		return manifest.Bytes(manifest.Primitive("var_u32"))
	case "nbt", "lnbt":
		return manifest.Primitive("nbt_le")
	case "bool", "boolean":
		return manifest.Primitive("bool")
	}
	if strings.HasPrefix(strings.ToLower(name), "var") {
		return primitive(name, nil, "integer")
	}
	known := map[string]string{"byte": "i8", "int8": "i8", "ubyte": "u8", "uint8": "u8", "short": "i16le", "int16": "i16le", "ushort": "u16le", "uint16": "u16le", "int": "i32le", "int32": "i32le", "uint": "u32le", "uint32": "u32le", "long": "i64le", "int64": "i64le", "ulong": "u64le", "uint64": "u64le", "float": "f32le", "float32": "f32le", "double": "f64le", "float64": "f64le", "uuid": "uuid"}
	if code, ok := known[strings.ToLower(name)]; ok {
		return manifest.Primitive(code)
	}
	return manifest.Unresolved("unknown Endstone scalar "+name, false)
}

func lowerEndstoneRepeat(repeat map[string]any, inner manifest.Node, context string) manifest.Node {
	if prefix := asString(repeat["prefix"]); prefix != "" {
		prefixNode := primitive(prefix, nil, "integer")
		if prefixNode.Kind == manifest.KindUnresolved {
			return manifest.Unresolved("unknown Endstone repeat prefix "+prefix+" at "+context, true)
		}
		return manifest.Array(prefixNode, inner)
	}
	if count, ok := asInt(repeat["count"]); ok && count > 0 {
		return manifest.FixedArray(uint64(count), inner)
	}
	return manifest.Unresolved("Endstone repeat lacks explicit prefix/count at "+context, true)
}

func minInt(left, right int) int {
	if left < right {
		return left
	}
	return right
}
