package gophertunneloracle

import (
	"fmt"
	"sort"
	"strings"

	"protocolgen/internal/manifest"
)

// shapeExpr is a compact wire-language expression. It keeps alternatives
// symbolic, so independent optional fields do not require a cartesian product
// merely to prove that both implementations describe the same paths.
type shapeExpr struct {
	kind   string
	token  string
	parts  []shapeExpr
	alts   []shapeExpr
	reason string
}

func emptyExpr() shapeExpr { return shapeExpr{kind: "empty"} }

func tokenExpr(token string) shapeExpr { return shapeExpr{kind: "token", token: token} }

func unknownExpr(reason string) shapeExpr { return shapeExpr{kind: "unknown", reason: reason} }

func concatExpr(expressions ...shapeExpr) shapeExpr {
	parts := make([]shapeExpr, 0, len(expressions))
	for _, expression := range expressions {
		if expression.kind == "unknown" {
			return expression
		}
		if expression.kind == "empty" {
			continue
		}
		if expression.kind == "concat" {
			parts = append(parts, expression.parts...)
			continue
		}
		parts = append(parts, expression)
	}
	switch len(parts) {
	case 0:
		return emptyExpr()
	case 1:
		return parts[0]
	default:
		return shapeExpr{kind: "concat", parts: parts}
	}
}

func altExpr(expressions ...shapeExpr) shapeExpr {
	seen := map[string]bool{}
	pending := append([]shapeExpr(nil), expressions...)
	alts := make([]shapeExpr, 0, len(pending))
	for len(pending) > 0 {
		expression := pending[0]
		pending = pending[1:]
		if expression.kind == "unknown" {
			return expression
		}
		if expression.kind == "alt" {
			pending = append(pending, expression.alts...)
			continue
		}
		key := expressionKey(expression)
		if seen[key] {
			continue
		}
		seen[key] = true
		alts = append(alts, expression)
	}
	sort.Slice(alts, func(i, j int) bool { return expressionKey(alts[i]) < expressionKey(alts[j]) })
	switch len(alts) {
	case 0:
		return emptyExpr()
	case 1:
		return alts[0]
	default:
		return shapeExpr{kind: "alt", alts: alts}
	}
}

func expressionKey(expression shapeExpr) string {
	switch expression.kind {
	case "empty":
		return "e"
	case "unknown":
		return "u:" + expression.reason
	case "token":
		return "t:" + expression.token
	case "concat":
		parts := make([]string, len(expression.parts))
		for index, part := range expression.parts {
			parts[index] = expressionKey(part)
		}
		return "c:[" + strings.Join(parts, ",") + "]"
	case "alt":
		alts := make([]string, len(expression.alts))
		for index, alternative := range expression.alts {
			alts[index] = expressionKey(alternative)
		}
		return "a:[" + strings.Join(alts, ",") + "]"
	default:
		return "?"
	}
}

func symbolicAgreement(packet manifest.Packet, source sourcePacket) bool {
	want := canonicalPacketExpr(packet)
	got := sourceSequenceExpr(source.Operations)
	if want.kind == "unknown" || got.kind == "unknown" {
		return false
	}
	return expressionKey(want) == expressionKey(got)
}

func canonicalPacketExpr(packet manifest.Packet) shapeExpr {
	fields := append([]manifest.Field(nil), packet.Fields...)
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
	parts := make([]shapeExpr, 0, len(fields))
	for _, field := range fields {
		parts = append(parts, canonicalNodeExpr(field.Name, field.Encode))
	}
	return concatExpr(parts...)
}

func canonicalNodeExpr(path string, node manifest.Node) shapeExpr {
	switch node.Kind {
	case manifest.KindVoid:
		return emptyExpr()
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return unknownExpr("manifest: primitive at " + path + " has no shape")
		}
		if node.Primitive.Code == "uuid" {
			return tokenExpr("UUID16")
		}
		return tokenExpr("P:" + canonicalPrimitive(node.Primitive.Code))
	case manifest.KindEnum:
		if node.Primitive == nil {
			return unknownExpr("manifest: enum at " + path + " has no underlying shape")
		}
		return tokenExpr("P:" + canonicalPrimitive(node.Primitive.Code))
	case manifest.KindString, manifest.KindBytes:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return unknownExpr("manifest: " + path + ": " + err.Error())
		}
		return tokenExpr("LEN:" + canonicalPrimitive(prefix))
	case manifest.KindBitset:
		if node.Length == 0 {
			return unknownExpr("manifest: bitset at " + path + " has no length")
		}
		return tokenExpr(fmt.Sprintf("BITSET:%d", node.Length))
	case manifest.KindArray:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil || node.Element == nil {
			if err == nil {
				err = fmt.Errorf("array has no element")
			}
			return unknownExpr("manifest: " + path + ": " + err.Error())
		}
		if isManifestU8(*node.Element) {
			return tokenExpr("LEN:" + canonicalPrimitive(prefix))
		}
		return concatExpr(tokenExpr("ARRAY:"+canonicalPrimitive(prefix)), canonicalNodeExpr(path+"[]", *node.Element), tokenExpr("/ARRAY"))
	case manifest.KindFixedArray:
		if node.Element == nil || node.Length == 0 {
			return unknownExpr("manifest: fixed array at " + path + " is incomplete")
		}
		if node.Length == 16 && isManifestU8(*node.Element) {
			return tokenExpr("UUID16")
		}
		return concatExpr(tokenExpr(fmt.Sprintf("FIXED:%d", node.Length)), canonicalNodeExpr(path+"[]", *node.Element), tokenExpr("/FIXED"))
	case manifest.KindSequence:
		parts := make([]shapeExpr, 0, len(node.Elements))
		for index, child := range node.Elements {
			parts = append(parts, canonicalNodeExpr(fmt.Sprintf("%s[%d]", path, index), child))
		}
		return concatExpr(parts...)
	case manifest.KindOptional:
		if node.Value == nil {
			return unknownExpr("manifest: optional at " + path + " has no value")
		}
		presence := tokenExpr("OPTION:bool")
		close := tokenExpr("/OPTION")
		return altExpr(concatExpr(presence, close), concatExpr(presence, canonicalNodeExpr(path, *node.Value), close))
	case manifest.KindStruct:
		fields := append([]manifest.Field(nil), node.Fields...)
		sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
		parts := make([]shapeExpr, 0, len(fields))
		for _, field := range fields {
			parts = append(parts, canonicalNodeExpr(path+"."+field.Name, field.Encode))
		}
		return concatExpr(parts...)
	case manifest.KindMap:
		if node.Prefix == nil || node.Key == nil || node.Value == nil {
			return unknownExpr("manifest: map at " + path + " is incomplete")
		}
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return unknownExpr("manifest: " + path + ": " + err.Error())
		}
		return concatExpr(tokenExpr("ARRAY:"+canonicalPrimitive(prefix)), canonicalNodeExpr(path+".<key>", *node.Key), canonicalNodeExpr(path+".<value>", *node.Value), tokenExpr("/ARRAY"))
	case manifest.KindUnion:
		control, err := manifestPrimitive(node.Control)
		if err != nil {
			return unknownExpr("manifest: union at " + path + ": " + err.Error())
		}
		branches := make([]shapeExpr, 0, len(node.Variants))
		for _, variant := range node.Variants {
			branches = append(branches, concatExpr(tokenExpr(fmt.Sprintf("VARIANT:%d", variant.Value)), canonicalNodeExpr(path+".variant", variant.Encode)))
		}
		if len(branches) == 0 {
			return unknownExpr("manifest: union at " + path + " has no variants")
		}
		return concatExpr(tokenExpr("P:"+canonicalPrimitive(control)), altExpr(branches...))
	case manifest.KindConditional:
		branches := make([]shapeExpr, 0, len(node.Cases)+1)
		for index, oneCase := range node.Cases {
			parts := make([]shapeExpr, 0, len(oneCase.Encode))
			for childIndex, child := range oneCase.Encode {
				parts = append(parts, canonicalNodeExpr(fmt.Sprintf("%s.case[%d].%d", path, index, childIndex), child))
			}
			branches = append(branches, concatExpr(parts...))
		}
		if node.Default != nil {
			branches = append(branches, canonicalNodeExpr(path+".default", *node.Default))
		}
		if len(branches) == 0 {
			return unknownExpr("manifest: conditional at " + path + " has no finite cases")
		}
		return altExpr(branches...)
	case manifest.KindReserved, manifest.KindIgnored:
		if node.Element == nil {
			return unknownExpr("manifest: compatibility node at " + path + " has no element")
		}
		return canonicalNodeExpr(path, *node.Element)
	case manifest.KindRecursive, manifest.KindOpaque, manifest.KindUnresolved:
		return unknownExpr("manifest: " + string(node.Kind) + " at " + path + ": " + node.Reason)
	default:
		return unknownExpr("manifest: unsupported node " + string(node.Kind) + " at " + path)
	}
}

func sourceSequenceExpr(operations []sourceOperation) shapeExpr {
	parts := make([]shapeExpr, 0, len(operations))
	for index, operation := range operations {
		if index+1 < len(operations) {
			next := operations[index+1]
			if next.ConsumesPrefix && operation.Kind == "primitive" && operation.Code == next.Prefix {
				continue
			}
		}
		parts = append(parts, sourceOperationExpr(operation))
	}
	return concatExpr(parts...)
}

func sourceOperationExpr(operation sourceOperation) shapeExpr {
	path := operation.Field
	switch operation.Kind {
	case "primitive":
		code := operation.Code
		if code == "raw_bytes" && isPreencodedNBTField(path) {
			code = "nbt_le"
		}
		return tokenExpr("P:" + canonicalPrimitive(code))
	case "string", "bytes":
		return tokenExpr("LEN:" + canonicalPrimitive(operation.Prefix))
	case "uuid":
		return tokenExpr("UUID16")
	case "bitset":
		if operation.Length == 0 {
			return unknownExpr("gophertunnel: bitset at " + path + " has no static length")
		}
		return tokenExpr(fmt.Sprintf("BITSET:%d", operation.Length))
	case "variant_marker":
		return tokenExpr(fmt.Sprintf("VARIANT:%d", operation.VariantValue))
	case "array":
		if isSourceU8(operation.Element) {
			return tokenExpr("LEN:" + canonicalPrimitive(operation.Prefix))
		}
		return concatExpr(tokenExpr("ARRAY:"+canonicalPrimitive(operation.Prefix)), sourceSequenceExpr(operation.Element), tokenExpr("/ARRAY"))
	case "fixed_array":
		if operation.Length == 0 {
			return unknownExpr("gophertunnel: fixed array at " + path + " has no length")
		}
		if operation.Length == 16 && isSourceU8(operation.Element) {
			return tokenExpr("UUID16")
		}
		return concatExpr(tokenExpr(fmt.Sprintf("FIXED:%d", operation.Length)), sourceSequenceExpr(operation.Element), tokenExpr("/FIXED"))
	case "optional":
		presence := tokenExpr("OPTION:" + canonicalPrimitive(operation.Presence))
		close := tokenExpr("/OPTION")
		return altExpr(concatExpr(presence, close), concatExpr(presence, sourceSequenceExpr(operation.Value), close))
	case "union":
		branches := make([]shapeExpr, 0, len(operation.Variants))
		for _, variant := range operation.Variants {
			branches = append(branches, concatExpr(tokenExpr(fmt.Sprintf("VARIANT:%d", variant.Value)), sourceSequenceExpr(variant.Ops)))
		}
		if len(branches) == 0 {
			return unknownExpr("gophertunnel: union at " + path + " has no statically known variants")
		}
		if operation.Control == "" {
			return altExpr(branches...)
		}
		return concatExpr(tokenExpr("P:"+canonicalPrimitive(operation.Control)), altExpr(branches...))
	case "conditional", "switch", "type_switch":
		branches := make([]shapeExpr, 0, len(operation.Variants)+1)
		for _, variant := range operation.Variants {
			values := variant.Values
			if len(values) == 0 {
				values = []int64{variant.Value}
			}
			for _, value := range values {
				branch := sourceSequenceExpr(variant.Ops)
				if operation.Kind == "switch" {
					branch = concatExpr(tokenExpr(fmt.Sprintf("VARIANT:%d", value)), branch)
				}
				branches = append(branches, branch)
			}
		}
		if operation.HasDefault {
			branches = append(branches, sourceSequenceExpr(operation.Default))
		}
		if len(branches) == 0 {
			return unknownExpr("gophertunnel: control-flow operation at " + path + " has no finite paths")
		}
		return altExpr(branches...)
	case "unresolved", "recursive":
		reason := operation.Reason
		if reason == "" {
			reason = "operation is not statically resolvable"
		}
		return unknownExpr("gophertunnel: " + reason + " at " + path)
	default:
		return unknownExpr("gophertunnel: unsupported operation " + operation.Kind + " at " + path)
	}
}
