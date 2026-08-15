package gophertunneloracle

import (
	"fmt"
	"sort"
	"strings"

	"protocolgen/internal/manifest"
)

// maxPathExpansion keeps a malicious or merely very union-heavy source from
// turning verification into an unbounded cartesian-product computation. The
// extractor reports the overflow as unresolved rather than silently checking a
// prefix of the possible wire paths.
const maxPathExpansion = 256

type wirePath struct {
	Constraint string
	Atoms      []atom
	Reasons    []string
}

type pathFragment struct {
	Constraints []string
	Operations  []sourceOperation
	Atoms       []atom
	Reasons     []string
}

func sourcePacketPaths(packet sourcePacket) []wirePath {
	if len(packet.Paths) == 0 {
		atoms, reasons := sourceAtoms(packet)
		return []wirePath{{Atoms: atoms, Reasons: reasons}}
	}
	paths := make([]wirePath, 0, len(packet.Paths))
	for _, path := range packet.Paths {
		atoms, reasons := sourcePathAtoms(path)
		paths = append(paths, wirePath{Constraint: strings.Join(path.Constraints, " && "), Atoms: atoms, Reasons: reasons})
	}
	return paths
}

func sourcePathAtoms(path sourcePath) ([]atom, []string) {
	var atoms []atom
	reasons := append([]string{}, path.Reasons...)
	for index, operation := range path.Operations {
		if index+1 < len(path.Operations) {
			next := path.Operations[index+1]
			if next.ConsumesPrefix && operation.Kind == "primitive" && operation.Code == next.Prefix {
				// The primitive is the explicit runtime length consumed by the
				// following array helper. Keep the length in the array shape once;
				// do not invent a second prefix or compare it as a payload value.
				continue
			}
		}
		one, unresolved := sourceWireOperationAtoms(operation)
		for index := range one {
			if one[index].Site == "" {
				one[index].Site = operation.Site
			}
		}
		atoms = append(atoms, one...)
		reasons = append(reasons, unresolved...)
	}
	return atoms, uniqueStrings(reasons)
}

func sourceWireOperationAtoms(operation sourceOperation) ([]atom, []string) {
	if operation.Kind == "union" && len(operation.Variants) == 1 {
		control := operation.Control
		var result []atom
		if control != "" {
			result = append(result, atom{Token: "P:" + canonicalPrimitive(control), Field: operation.Field, Display: control})
		}
		result = append(result, atom{Token: fmt.Sprintf("VARIANT:%d", operation.Variants[0].Value), Field: operation.Field, Display: fmt.Sprintf("variant(%d)", operation.Variants[0].Value)})
		for _, child := range operation.Variants[0].Ops {
			atoms, reasons := sourceWireOperationAtoms(child)
			result = append(result, atoms...)
			if len(reasons) > 0 {
				return result, reasons
			}
		}
		return result, nil
	}
	if operation.Kind == "conditional" || operation.Kind == "switch" || operation.Kind == "type_switch" {
		return nil, []string{"gophertunnel: unexpanded control-flow operation at " + operation.Field}
	}
	if operation.Kind == "array" || operation.Kind == "fixed_array" {
		var result []atom
		var reasons []string
		if operation.Kind == "array" {
			if isSourceU8(operation.Element) {
				return []atom{{Token: "LEN:" + canonicalPrimitive(operation.Prefix), Field: operation.Field, Display: "byte-array(prefix=" + operation.Prefix + ")"}}, nil
			}
			result = append(result, atom{Token: "ARRAY:" + canonicalPrimitive(operation.Prefix), Field: operation.Field, Display: "array(prefix=" + operation.Prefix + ")"})
		} else {
			if operation.Length == 16 && isSourceU8(operation.Element) {
				return []atom{{Token: "UUID16", Field: operation.Field, Display: "uuid(16 bytes)"}}, nil
			}
			result = append(result, atom{Token: fmt.Sprintf("FIXED:%d", operation.Length), Field: operation.Field, Display: fmt.Sprintf("fixed-array(length=%d)", operation.Length)})
		}
		for _, child := range operation.Element {
			atoms, childReasons := sourceWireOperationAtoms(child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		if operation.Kind == "array" {
			result = append(result, atom{Token: "/ARRAY", Field: operation.Field, Display: "/array"})
		} else {
			result = append(result, atom{Token: "/FIXED", Field: operation.Field, Display: "/fixed-array"})
		}
		return result, uniqueStrings(reasons)
	}
	if operation.Kind == "optional" {
		result := []atom{{Token: "OPTION:" + canonicalPrimitive(operation.Presence), Field: operation.Field, Display: "option(presence=" + operation.Presence + ")"}}
		var reasons []string
		for _, child := range operation.Value {
			atoms, childReasons := sourceWireOperationAtoms(child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		result = append(result, atom{Token: "/OPTION", Field: operation.Field, Display: "/option"})
		return result, uniqueStrings(reasons)
	}
	return sourceOperationAtoms(operation)
}

// expandSourcePaths expands only finite control flow. Arrays remain arrays,
// but their element alternatives are expanded so a union nested inside an
// element is compared as a union variant rather than being flattened into a
// guessed single sequence.
func expandSourcePaths(operations []sourceOperation) []sourcePath {
	fragments := []pathFragment{{}}
	for _, operation := range operations {
		alternatives := expandSourceOperation(operation)
		fragments = combinePathFragments(fragments, alternatives)
		if len(fragments) == 1 && len(fragments[0].Reasons) > 0 {
			return []sourcePath{{Operations: fragments[0].Operations, Reasons: fragments[0].Reasons}}
		}
	}
	paths := make([]sourcePath, 0, len(fragments))
	for _, fragment := range fragments {
		paths = append(paths, sourcePath{Constraints: fragment.Constraints, Operations: fragment.Operations, Reasons: fragment.Reasons})
	}
	return paths
}

func expandSourceOperation(operation sourceOperation) []pathFragment {
	switch operation.Kind {
	case "conditional", "switch", "type_switch":
		var result []pathFragment
		for _, variant := range operation.Variants {
			children := expandSourcePaths(variant.Ops)
			if len(children) == 0 {
				children = []sourcePath{{}}
			}
			values := variant.Values
			if len(values) == 0 {
				values = []int64{variant.Value}
			}
			for _, value := range values {
				for _, child := range children {
					operations := append([]sourceOperation(nil), child.Operations...)
					if operation.Kind == "switch" {
						operations = append([]sourceOperation{{Kind: "variant_marker", Field: operation.Field, VariantValue: value, Site: operation.Site}}, operations...)
					}
					constraints := append([]string{}, child.Constraints...)
					constraint := variant.Constraint
					if constraint == "" && operation.CompareTo != "" && operation.Kind == "switch" {
						constraint = fmt.Sprintf("%s=%d", operation.CompareTo, value)
					}
					if constraint == "" {
						constraint = variant.Name
					}
					if constraint != "" {
						constraints = append([]string{constraint}, constraints...)
					}
					result = append(result, pathFragment{Constraints: constraints, Operations: operations})
				}
			}
		}
		if operation.HasDefault || len(operation.Variants) == 0 {
			children := expandSourcePaths(operation.Default)
			if len(children) == 0 {
				children = []sourcePath{{}}
			}
			for _, child := range children {
				constraints := append([]string{"default"}, child.Constraints...)
				result = append(result, pathFragment{Constraints: constraints, Operations: child.Operations})
			}
		}
		if len(result) == 0 {
			return []pathFragment{{Operations: []sourceOperation{{Kind: "unresolved", Field: operation.Field, Reason: "control-flow statement has no finite path"}}}}
		}
		return result
	case "union":
		var result []pathFragment
		for _, variant := range operation.Variants {
			children := expandSourcePaths(variant.Ops)
			if len(children) == 0 {
				children = []sourcePath{{}}
			}
			values := variant.Values
			if len(values) == 0 {
				values = []int64{variant.Value}
			}
			for _, value := range values {
				for _, child := range children {
					one := sourceOperation{Kind: "union", Field: operation.Field, Control: operation.Control, Variants: []sourceVariant{{
						Value: value, Name: variant.Name, Ops: child.Operations,
					}}}
					constraints := append([]string{}, child.Constraints...)
					if variant.Constraint != "" {
						constraints = append([]string{variant.Constraint}, constraints...)
					}
					result = append(result, pathFragment{Constraints: constraints, Operations: []sourceOperation{one}})
				}
			}
		}
		if len(result) == 0 {
			return []pathFragment{{Operations: []sourceOperation{{Kind: "unresolved", Field: operation.Field, Reason: "union has no statically known variants"}}}}
		}
		return result
	case "optional":
		absent := operation
		absent.Value = nil
		present := operation
		children := expandSourcePaths(operation.Value)
		if len(children) == 0 {
			children = []sourcePath{{}}
		}
		result := []pathFragment{{Constraints: []string{operation.Field + " absent"}, Operations: []sourceOperation{absent}}}
		for _, child := range children {
			present.Value = child.Operations
			constraints := append([]string{operation.Field + " present"}, child.Constraints...)
			result = append(result, pathFragment{Constraints: constraints, Operations: []sourceOperation{present}})
		}
		return result
	case "array", "fixed_array":
		children := expandSourcePaths(operation.Element)
		if len(children) == 0 {
			children = []sourcePath{{}}
		}
		result := make([]pathFragment, 0, len(children))
		for _, child := range children {
			one := operation
			one.Element = child.Operations
			result = append(result, pathFragment{Constraints: child.Constraints, Operations: []sourceOperation{one}})
		}
		return result
	default:
		return []pathFragment{{Operations: []sourceOperation{operation}}}
	}
}

func combinePathFragments(left, right []pathFragment) []pathFragment {
	if len(left) == 0 || len(right) == 0 {
		return nil
	}
	if len(left) > maxPathExpansion/len(right) {
		return []pathFragment{{Reasons: []string{fmt.Sprintf("control-flow path expansion exceeds limit %d", maxPathExpansion)}}}
	}
	result := make([]pathFragment, 0, len(left)*len(right))
	for _, l := range left {
		for _, r := range right {
			constraints := append(append([]string{}, l.Constraints...), r.Constraints...)
			operations := append(append([]sourceOperation{}, l.Operations...), r.Operations...)
			atoms := append(append([]atom{}, l.Atoms...), r.Atoms...)
			reasons := append(append([]string{}, l.Reasons...), r.Reasons...)
			result = append(result, pathFragment{Constraints: constraints, Operations: operations, Atoms: atoms, Reasons: reasons})
		}
	}
	return result
}

func canonicalPacketPaths(packet manifest.Packet) []wirePath {
	fragments := []pathFragment{{}}
	fields := append([]manifest.Field(nil), packet.Fields...)
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
	for _, field := range fields {
		alternatives := canonicalNodePaths(field.Name, field.Encode)
		fragments = combinePathFragments(fragments, alternatives)
	}
	paths := make([]wirePath, 0, len(fragments))
	for _, fragment := range fragments {
		paths = append(paths, wirePath{Constraint: strings.Join(fragment.Constraints, " && "), Atoms: fragment.Atoms, Reasons: uniqueStrings(fragment.Reasons)})
	}
	return paths
}

func canonicalNodePaths(path string, node manifest.Node) []pathFragment {
	single := func(atoms ...atom) []pathFragment { return []pathFragment{{Atoms: atoms}} }
	switch node.Kind {
	case manifest.KindVoid:
		return []pathFragment{{}}
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return []pathFragment{{Reasons: []string{"manifest: primitive at " + path + " has no shape"}}}
		}
		if node.Primitive.Code == "uuid" {
			return single(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"})
		}
		return single(atom{Token: "P:" + canonicalPrimitive(node.Primitive.Code), Field: path, Display: node.Primitive.Code})
	case manifest.KindEnum:
		if node.Primitive == nil {
			return []pathFragment{{Reasons: []string{"manifest: enum at " + path + " has no underlying shape"}}}
		}
		return single(atom{Token: "P:" + canonicalPrimitive(node.Primitive.Code), Field: path, Display: "enum(" + node.Primitive.Code + ")"})
	case manifest.KindString, manifest.KindBytes:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return []pathFragment{{Reasons: []string{"manifest: " + path + ": " + err.Error()}}}
		}
		kind := "string"
		if node.Kind == manifest.KindBytes {
			kind = "bytes"
		}
		return single(atom{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: kind + "(prefix=" + prefix + ")"})
	case manifest.KindBitset:
		if node.Length == 0 {
			return []pathFragment{{Reasons: []string{"manifest: bitset at " + path + " has no length"}}}
		}
		return single(atom{Token: fmt.Sprintf("BITSET:%d", node.Length), Field: path, Display: fmt.Sprintf("bitset(length=%d)", node.Length)})
	case manifest.KindArray:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil || node.Element == nil {
			if err == nil {
				err = fmt.Errorf("array has no element")
			}
			return []pathFragment{{Reasons: []string{"manifest: " + path + ": " + err.Error()}}}
		}
		if isManifestU8(*node.Element) {
			return single(atom{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: "byte-array(prefix=" + prefix + ")"})
		}
		children := canonicalNodePaths(path+"[]", *node.Element)
		for i := range children {
			children[i].Atoms = append([]atom{{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "array(prefix=" + prefix + ")"}}, children[i].Atoms...)
			children[i].Atoms = append(children[i].Atoms, atom{Token: "/ARRAY", Field: path, Display: "/array"})
		}
		return children
	case manifest.KindFixedArray:
		if node.Element == nil || node.Length == 0 {
			return []pathFragment{{Reasons: []string{"manifest: fixed array at " + path + " is incomplete"}}}
		}
		if node.Length == 16 && isManifestU8(*node.Element) {
			return single(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"})
		}
		children := canonicalNodePaths(path+"[]", *node.Element)
		for i := range children {
			children[i].Atoms = append([]atom{{Token: fmt.Sprintf("FIXED:%d", node.Length), Field: path, Display: fmt.Sprintf("fixed-array(length=%d)", node.Length)}}, children[i].Atoms...)
			children[i].Atoms = append(children[i].Atoms, atom{Token: "/FIXED", Field: path, Display: "/fixed-array"})
		}
		return children
	case manifest.KindSequence:
		return canonicalSequencePaths(path, node.Elements)
	case manifest.KindOptional:
		if node.Value == nil {
			return []pathFragment{{Reasons: []string{"manifest: optional at " + path + " has no value"}}}
		}
		absent := pathFragment{Constraints: []string{path + " absent"}, Atoms: []atom{{Token: "OPTION:bool", Field: path, Display: "option(presence=bool)"}, {Token: "/OPTION", Field: path, Display: "/option"}}}
		present := canonicalNodePaths(path, *node.Value)
		result := make([]pathFragment, 0, len(present)+1)
		result = append(result, absent)
		for _, child := range present {
			child.Constraints = append([]string{path + " present"}, child.Constraints...)
			child.Atoms = append([]atom{{Token: "OPTION:bool", Field: path, Display: "option(presence=bool)"}}, child.Atoms...)
			child.Atoms = append(child.Atoms, atom{Token: "/OPTION", Field: path, Display: "/option"})
			result = append(result, child)
		}
		return result
	case manifest.KindStruct:
		fields := append([]manifest.Field(nil), node.Fields...)
		sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
		var result []pathFragment
		result = []pathFragment{{}}
		for _, field := range fields {
			result = combinePathFragments(result, canonicalNodePaths(path+"."+field.Name, field.Encode))
		}
		return result
	case manifest.KindMap:
		if node.Prefix == nil || node.Key == nil || node.Value == nil {
			return []pathFragment{{Reasons: []string{"manifest: map at " + path + " is incomplete"}}}
		}
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return []pathFragment{{Reasons: []string{"manifest: " + path + ": " + err.Error()}}}
		}
		children := combinePathFragments(canonicalNodePaths(path+".<key>", *node.Key), canonicalNodePaths(path+".<value>", *node.Value))
		for i := range children {
			children[i].Atoms = append([]atom{{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "map(prefix=" + prefix + ")"}}, children[i].Atoms...)
			children[i].Atoms = append(children[i].Atoms, atom{Token: "/ARRAY", Field: path, Display: "/map"})
		}
		return children
	case manifest.KindUnion:
		control, err := manifestPrimitive(node.Control)
		if err != nil {
			return []pathFragment{{Reasons: []string{"manifest: union at " + path + ": " + err.Error()}}}
		}
		variants := append([]manifest.Variant(nil), node.Variants...)
		sort.SliceStable(variants, func(i, j int) bool { return variants[i].Value < variants[j].Value })
		var result []pathFragment
		for _, variant := range variants {
			children := canonicalNodePaths(path+".variant", variant.Encode)
			for _, child := range children {
				child.Constraints = append([]string{fmt.Sprintf("%s variant=%d", path, variant.Value)}, child.Constraints...)
				prefixAtoms := []atom{
					{Token: "P:" + canonicalPrimitive(control), Field: path, Display: control},
					{Token: fmt.Sprintf("VARIANT:%d", variant.Value), Field: path, Display: fmt.Sprintf("variant(%d)", variant.Value)},
				}
				child.Atoms = append(prefixAtoms, child.Atoms...)
				result = append(result, child)
			}
		}
		if len(result) == 0 {
			return []pathFragment{{Reasons: []string{"manifest: union at " + path + " has no variants"}}}
		}
		return result
	case manifest.KindConditional:
		var result []pathFragment
		for _, oneCase := range node.Cases {
			children := canonicalSequencePaths(path, oneCase.Encode)
			for _, child := range children {
				child.Constraints = append([]string{fmt.Sprintf("%s case=%s", path, oneCase.Value)}, child.Constraints...)
				result = append(result, child)
			}
		}
		if node.Default != nil {
			children := canonicalNodePaths(path+".default", *node.Default)
			for _, child := range children {
				child.Constraints = append([]string{path + " default"}, child.Constraints...)
				result = append(result, child)
			}
		}
		if len(result) == 0 {
			return []pathFragment{{Reasons: []string{"manifest: conditional at " + path + " has no finite cases"}}}
		}
		return result
	case manifest.KindReserved, manifest.KindIgnored:
		if node.Element == nil {
			return []pathFragment{{Reasons: []string{"manifest: compatibility node at " + path + " has no element"}}}
		}
		return canonicalNodePaths(path, *node.Element)
	case manifest.KindRecursive:
		return []pathFragment{{Reasons: []string{"manifest: recursive node at " + path + " is not statically finite"}}}
	case manifest.KindOpaque, manifest.KindUnresolved:
		return []pathFragment{{Reasons: []string{"manifest: " + string(node.Kind) + " at " + path + ": " + node.Reason}}}
	default:
		return []pathFragment{{Reasons: []string{"manifest: unsupported node " + string(node.Kind) + " at " + path}}}
	}
}

func canonicalSequencePaths(path string, nodes []manifest.Node) []pathFragment {
	result := []pathFragment{{}}
	for index, node := range nodes {
		result = combinePathFragments(result, canonicalNodePaths(fmt.Sprintf("%s[%d]", path, index), node))
	}
	return result
}
