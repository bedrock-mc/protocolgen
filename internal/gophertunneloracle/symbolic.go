package gophertunneloracle

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"protocolgen/internal/manifest"
)

// shapeExpr is a wire-language expression: a finite regular language over
// wire atoms in which alternatives stay symbolic, so independent optional
// fields never need a cartesian product to be compared.
type shapeExpr struct {
	kind    string
	atom    atom
	parts   []shapeExpr
	alts    []shapeExpr
	reasons []string
}

// recursionUnroll bounds how many times a recursive type is expanded on both
// sides before the remaining recursion is compared as one RECURSE atom.
const recursionUnroll = 2

func emptyExpr() shapeExpr { return shapeExpr{kind: "empty"} }

func atomExpr(a atom) shapeExpr { return shapeExpr{kind: "token", atom: a} }

func unknownExpr(reasons ...string) shapeExpr { return shapeExpr{kind: "unknown", reasons: reasons} }

func recurseExpr(field string) shapeExpr {
	return atomExpr(atom{Token: "RECURSE", Field: field, Display: fmt.Sprintf("recurse(depth>%d)", recursionUnroll)})
}

// concatExpr sequences expressions; an unknown part makes the whole sequence
// unknown while still collecting every reason.
func concatExpr(expressions ...shapeExpr) shapeExpr {
	var reasons []string
	parts := make([]shapeExpr, 0, len(expressions))
	for _, expression := range expressions {
		switch expression.kind {
		case "unknown":
			reasons = append(reasons, expression.reasons...)
		case "empty":
		case "concat":
			parts = append(parts, expression.parts...)
		default:
			parts = append(parts, expression)
		}
	}
	if len(reasons) > 0 {
		return unknownExpr(uniqueStrings(reasons)...)
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
	var reasons []string
	seen := map[string]bool{}
	pending := append([]shapeExpr(nil), expressions...)
	alts := make([]shapeExpr, 0, len(pending))
	for len(pending) > 0 {
		expression := pending[0]
		pending = pending[1:]
		if expression.kind == "unknown" {
			reasons = append(reasons, expression.reasons...)
			continue
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
	if len(reasons) > 0 {
		return unknownExpr(uniqueStrings(reasons)...)
	}
	sort.SliceStable(alts, func(i, j int) bool { return expressionKey(alts[i]) < expressionKey(alts[j]) })
	switch len(alts) {
	case 0:
		return emptyExpr()
	case 1:
		return alts[0]
	default:
		return shapeExpr{kind: "alt", alts: alts}
	}
}

func repeatExpr(expression shapeExpr, count uint64) shapeExpr {
	parts := make([]shapeExpr, 0, count)
	for i := uint64(0); i < count; i++ {
		parts = append(parts, expression)
	}
	return concatExpr(parts...)
}

// optionalExpr is the wire shape shared by manifest optionals and bool-guarded
// gophertunnel fields: the presence flag followed by the value or nothing.
func optionalExpr(presence atom, value shapeExpr) shapeExpr {
	return concatExpr(atomExpr(presence), altExpr(emptyExpr(), value))
}

func expressionKey(expression shapeExpr) string {
	switch expression.kind {
	case "empty":
		return "e"
	case "unknown":
		return "u:" + strings.Join(expression.reasons, ";")
	case "token":
		return "t:" + expression.atom.Token
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

func primitiveAtom(code, field, display string) atom {
	return atom{Token: "P:" + canonicalPrimitive(code), Field: field, Display: display}
}

func variantAtom(value int64, field string) atom {
	return atom{Token: fmt.Sprintf("VARIANT:%d", value), Field: field, Display: fmt.Sprintf("variant(%d)", value)}
}

func canonicalPacketExpr(packet manifest.Packet) shapeExpr {
	fields := append([]manifest.Field(nil), packet.Fields...)
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
	ctx := &canonicalContext{ancestors: map[string]manifest.Node{}, depth: map[string]int{}}
	parts := make([]shapeExpr, 0, len(fields))
	for _, field := range fields {
		parts = append(parts, ctx.node(field.Name, field.Encode))
	}
	return concatExpr(parts...)
}

// canonicalContext tracks named ancestors so recursive nodes can be unrolled
// a bounded number of times.
type canonicalContext struct {
	ancestors map[string]manifest.Node
	depth     map[string]int
}

func (ctx *canonicalContext) node(path string, node manifest.Node) shapeExpr {
	if node.TypeID != "" {
		if _, seen := ctx.ancestors[node.TypeID]; !seen {
			ctx.ancestors[node.TypeID] = node
			defer delete(ctx.ancestors, node.TypeID)
		}
	}
	switch node.Kind {
	case manifest.KindVoid:
		return emptyExpr()
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return unknownExpr("manifest: primitive at " + path + " has no shape")
		}
		if node.Primitive.Code == "uuid" {
			return atomExpr(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"})
		}
		return atomExpr(primitiveAtom(node.Primitive.Code, path, node.Primitive.Code))
	case manifest.KindEnum:
		if node.Primitive == nil {
			return unknownExpr("manifest: enum at " + path + " has no underlying shape")
		}
		return atomExpr(primitiveAtom(node.Primitive.Code, path, "enum("+node.Primitive.Code+")"))
	case manifest.KindString, manifest.KindBytes:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return unknownExpr("manifest: " + path + ": " + err.Error())
		}
		kind := "string"
		if node.Kind == manifest.KindBytes {
			kind = "bytes"
		}
		return atomExpr(atom{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: kind + "(prefix=" + prefix + ")"})
	case manifest.KindBitset:
		if node.Length == 0 {
			return unknownExpr("manifest: bitset at " + path + " has no length")
		}
		return atomExpr(atom{Token: fmt.Sprintf("BITSET:%d", node.Length), Field: path, Display: fmt.Sprintf("bitset(length=%d)", node.Length)})
	case manifest.KindArray:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil || node.Element == nil {
			if err == nil {
				err = fmt.Errorf("array has no element")
			}
			return unknownExpr("manifest: " + path + ": " + err.Error())
		}
		if isManifestU8(*node.Element) {
			return atomExpr(atom{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: "byte-array(prefix=" + prefix + ")"})
		}
		return concatExpr(
			atomExpr(atom{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "array(prefix=" + prefix + ")"}),
			ctx.node(path+"[]", *node.Element),
			atomExpr(atom{Token: "/ARRAY", Field: path, Display: "/array"}),
		)
	case manifest.KindFixedArray:
		if node.Element == nil || node.Length == 0 {
			return unknownExpr("manifest: fixed array at " + path + " is incomplete")
		}
		if node.Length == 16 && isManifestU8(*node.Element) {
			return atomExpr(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"})
		}
		return repeatExpr(ctx.node(path+"[]", *node.Element), node.Length)
	case manifest.KindSequence:
		parts := make([]shapeExpr, 0, len(node.Elements))
		for index, child := range node.Elements {
			parts = append(parts, ctx.node(fmt.Sprintf("%s[%d]", path, index), child))
		}
		return concatExpr(parts...)
	case manifest.KindOptional:
		if node.Value == nil {
			return unknownExpr("manifest: optional at " + path + " has no value")
		}
		return optionalExpr(primitiveAtom("bool", path, "option(presence=bool)"), ctx.node(path, *node.Value))
	case manifest.KindStruct:
		fields := append([]manifest.Field(nil), node.Fields...)
		sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
		parts := make([]shapeExpr, 0, len(fields))
		for _, field := range fields {
			parts = append(parts, ctx.node(path+"."+field.Name, field.Encode))
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
		return concatExpr(
			atomExpr(atom{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "map(prefix=" + prefix + ")"}),
			ctx.node(path+".<key>", *node.Key),
			ctx.node(path+".<value>", *node.Value),
			atomExpr(atom{Token: "/ARRAY", Field: path, Display: "/map"}),
		)
	case manifest.KindUnion:
		control, err := manifestPrimitive(node.Control)
		if err != nil {
			return unknownExpr("manifest: union at " + path + ": " + err.Error())
		}
		branches := make([]shapeExpr, 0, len(node.Variants))
		for _, variant := range node.Variants {
			branches = append(branches, concatExpr(atomExpr(variantAtom(variant.Value, path)), ctx.node(path+".variant", variant.Encode)))
		}
		if len(branches) == 0 {
			return unknownExpr("manifest: union at " + path + " has no variants")
		}
		return concatExpr(atomExpr(primitiveAtom(control, path, "union(control="+control+")")), altExpr(branches...))
	case manifest.KindConditional:
		branches := make([]shapeExpr, 0, len(node.Cases)+1)
		for index, oneCase := range node.Cases {
			parts := make([]shapeExpr, 0, len(oneCase.Encode))
			for childIndex, child := range oneCase.Encode {
				parts = append(parts, ctx.node(fmt.Sprintf("%s.case[%d].%d", path, index, childIndex), child))
			}
			branches = append(branches, concatExpr(parts...))
		}
		if node.Default != nil {
			branches = append(branches, ctx.node(path+".default", *node.Default))
		}
		if len(branches) == 0 {
			return unknownExpr("manifest: conditional at " + path + " has no finite cases")
		}
		return altExpr(branches...)
	case manifest.KindReserved, manifest.KindIgnored:
		if node.Element == nil {
			return unknownExpr("manifest: compatibility node at " + path + " has no element")
		}
		return ctx.node(path, *node.Element)
	case manifest.KindRecursive:
		target, ok := ctx.ancestors[node.Target]
		if !ok {
			return unknownExpr("manifest: recursive node at " + path + " targets unknown type " + node.Target)
		}
		if ctx.depth[node.Target] >= recursionUnroll {
			return recurseExpr(path)
		}
		ctx.depth[node.Target]++
		defer func() { ctx.depth[node.Target]-- }()
		return ctx.node(path, target)
	case manifest.KindOpaque, manifest.KindUnresolved:
		return unknownExpr("manifest: " + string(node.Kind) + " at " + path + ": " + node.Reason)
	default:
		return unknownExpr("manifest: unsupported node " + string(node.Kind) + " at " + path)
	}
}

// sourceSequenceExpr builds the wire language of a gophertunnel operation
// sequence. Control flow that tests a field read earlier in the same sequence
// is hoisted next to that read, so Go code such as `if x.Type == 3` placed
// after other fields compares as the union variant the manifest annotates on
// the control read.
func sourceSequenceExpr(operations []sourceOperation) shapeExpr {
	operations = coalesceLengthPrefixes(operations)
	if control, ok := earliestControl(operations); ok {
		return hoistedExpr(operations, control)
	}
	parts := make([]shapeExpr, 0, len(operations))
	for _, operation := range operations {
		parts = append(parts, sourceOperationExpr(operation))
	}
	return concatExpr(parts...)
}

// earliestControl finds the first primitive read that a later discriminated
// conditional or switch in the same sequence tests.
func earliestControl(operations []sourceOperation) (int, bool) {
	for index, operation := range operations {
		if operation.Kind != "primitive" || operation.Field == "" {
			continue
		}
		for _, later := range operations[index+1:] {
			if testsField(later, operation.Field) {
				return index, true
			}
		}
	}
	return 0, false
}

func testsField(operation sourceOperation, field string) bool {
	return (operation.Kind == "conditional" || operation.Kind == "switch") && operation.CompareTo == field && hasDiscriminants(operation)
}

// hoistedExpr branches on every discriminant the sequence tests for the
// control read, resolving each dependent conditional for that value. The
// wildcard branch covers every discriminant no conditional handles.
func hoistedExpr(operations []sourceOperation, control int) shapeExpr {
	field := operations[control].Field
	tail := operations[control+1:]
	var values []int64
	seen := map[int64]bool{}
	for _, operation := range tail {
		if !testsField(operation, field) {
			continue
		}
		for _, variant := range operation.Variants {
			for _, value := range variantValues(variant) {
				if !seen[value] {
					seen[value] = true
					values = append(values, value)
				}
			}
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	handled := make([]string, 0, len(values))
	branches := make([]shapeExpr, 0, len(values)+1)
	for _, value := range values {
		handled = append(handled, strconv.FormatInt(value, 10))
		branches = append(branches, concatExpr(atomExpr(variantAtom(value, field)), sourceSequenceExpr(resolveDiscriminant(tail, field, value, false))))
	}
	if allowsOtherDiscriminants(tail, field) {
		branches = append(branches, concatExpr(atomExpr(wildcardAtom(handled, field)), sourceSequenceExpr(resolveDiscriminant(tail, field, 0, true))))
	}
	return concatExpr(sourceSequenceExpr(operations[:control+1]), altExpr(branches...))
}

// allowsOtherDiscriminants reports whether an unhandled discriminant still
// produces a wire path. An `if` falls through; a switch does only when it has
// a default that writes, since an invalid-input default is not a wire path.
func allowsOtherDiscriminants(operations []sourceOperation, field string) bool {
	for _, operation := range operations {
		if !testsField(operation, field) {
			continue
		}
		if operation.Kind == "switch" && !operation.HasDefault {
			return false
		}
	}
	return true
}

func variantValues(variant sourceVariant) []int64 {
	if len(variant.Values) > 0 {
		return variant.Values
	}
	return []int64{variant.Value}
}

// resolveDiscriminant replaces every conditional on field with the branch
// taken for value (or for any unhandled value when wildcard is set).
func resolveDiscriminant(operations []sourceOperation, field string, value int64, wildcard bool) []sourceOperation {
	var result []sourceOperation
	for _, operation := range operations {
		if !testsField(operation, field) {
			result = append(result, operation)
			continue
		}
		var chosen []sourceOperation
		matched := false
		for _, variant := range operation.Variants {
			contains := false
			for _, candidate := range variantValues(variant) {
				if !wildcard && candidate == value {
					contains = true
				}
			}
			if contains != variant.Negated {
				chosen, matched = variant.Ops, true
				break
			}
		}
		if !matched && operation.HasDefault {
			chosen = operation.Default
		}
		result = append(result, resolveDiscriminant(chosen, field, value, wildcard)...)
	}
	return result
}

// coalesceLengthPrefixes drops a primitive that is only the explicit runtime
// length consumed by the following array helper.
func coalesceLengthPrefixes(operations []sourceOperation) []sourceOperation {
	result := make([]sourceOperation, 0, len(operations))
	for index, operation := range operations {
		if index+1 < len(operations) {
			next := operations[index+1]
			if next.ConsumesPrefix && operation.Kind == "primitive" && operation.Code == next.Prefix {
				continue
			}
		}
		result = append(result, operation)
	}
	return result
}

// hasDiscriminants reports whether every branch of a control-flow operation
// is keyed by resolved integer discriminants. A bool guard is not one.
func hasDiscriminants(operation sourceOperation) bool {
	for _, variant := range operation.Variants {
		if operation.Kind == "conditional" && !variant.Discriminant {
			return false
		}
		if len(variant.Values) == 0 && !variant.Discriminant {
			return false
		}
	}
	return len(operation.Variants) > 0
}

type discriminantBranch struct {
	marker atom
	ops    []sourceOperation
}

// discriminantBranches lists the variant markers of a control-flow operation.
// A default (or negated) branch becomes a wildcard marker that matches every
// discriminant the explicit branches do not handle.
func discriminantBranches(operation sourceOperation) []discriminantBranch {
	var result []discriminantBranch
	var handled []string
	var negated []sourceOperation
	hasNegated := false
	for _, variant := range operation.Variants {
		values := variant.Values
		if len(values) == 0 {
			values = []int64{variant.Value}
		}
		for _, value := range values {
			handled = append(handled, strconv.FormatInt(value, 10))
			if variant.Negated {
				hasNegated = true
				negated = variant.Ops
				continue
			}
			result = append(result, discriminantBranch{marker: variantAtom(value, operation.Field), ops: variant.Ops})
		}
	}
	if hasNegated {
		// `x != v { A } else { B }`: A is every other discriminant, B is v.
		result = append(result, discriminantBranch{marker: wildcardAtom(handled, operation.Field), ops: negated})
		if operation.HasDefault {
			for _, value := range handled {
				parsed, _ := strconv.ParseInt(value, 10, 64)
				result = append(result, discriminantBranch{marker: variantAtom(parsed, operation.Field), ops: operation.Default})
			}
		}
		return result
	}
	if operation.HasDefault {
		result = append(result, discriminantBranch{marker: wildcardAtom(handled, operation.Field), ops: operation.Default})
	}
	return result
}

func wildcardAtom(handled []string, field string) atom {
	sort.Strings(handled)
	return atom{Token: wildcardVariantPrefix + strings.Join(handled, ","), Field: field, Display: "variant(not " + strings.Join(handled, ",") + ")"}
}

func sourceOperationExpr(operation sourceOperation) shapeExpr {
	path := operation.Field
	switch operation.Kind {
	case "primitive":
		code := operation.Code
		if code == "raw_bytes" && isPreencodedNBTField(path) {
			code = "nbt_le"
		}
		return atomExpr(primitiveAtom(code, path, code))
	case "string", "bytes":
		return atomExpr(atom{Token: "LEN:" + canonicalPrimitive(operation.Prefix), Field: path, Display: operation.Kind + "(prefix=" + operation.Prefix + ")", Site: operation.Site})
	case "uuid":
		return atomExpr(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)", Site: operation.Site})
	case "bitset":
		if operation.Length == 0 {
			return unknownExpr("gophertunnel: bitset at " + path + " has no static length")
		}
		return atomExpr(atom{Token: fmt.Sprintf("BITSET:%d", operation.Length), Field: path, Display: fmt.Sprintf("bitset(length=%d)", operation.Length), Site: operation.Site})
	case "variant_marker":
		return atomExpr(variantAtom(operation.VariantValue, path))
	case "array":
		if isSourceU8(operation.Element) {
			return atomExpr(atom{Token: "LEN:" + canonicalPrimitive(operation.Prefix), Field: path, Display: "byte-array(prefix=" + operation.Prefix + ")", Site: operation.Site})
		}
		return concatExpr(
			atomExpr(atom{Token: "ARRAY:" + canonicalPrimitive(operation.Prefix), Field: path, Display: "array(prefix=" + operation.Prefix + ")", Site: operation.Site}),
			sourceSequenceExpr(operation.Element),
			atomExpr(atom{Token: "/ARRAY", Field: path, Display: "/array"}),
		)
	case "fixed_array":
		if operation.Length == 0 {
			return unknownExpr("gophertunnel: fixed array at " + path + " has no length")
		}
		if operation.Length == 16 && isSourceU8(operation.Element) {
			return atomExpr(atom{Token: "UUID16", Field: path, Display: "uuid(16 bytes)", Site: operation.Site})
		}
		return repeatExpr(sourceSequenceExpr(operation.Element), operation.Length)
	case "optional":
		presence := primitiveAtom(operation.Presence, path, "option(presence="+operation.Presence+")")
		presence.Site = operation.Site
		return optionalExpr(presence, sourceSequenceExpr(operation.Value))
	case "union":
		branches := make([]shapeExpr, 0, len(operation.Variants))
		for _, variant := range operation.Variants {
			branches = append(branches, concatExpr(atomExpr(variantAtom(variant.Value, path)), sourceSequenceExpr(variant.Ops)))
		}
		if len(branches) == 0 {
			return unknownExpr("gophertunnel: union at " + path + " has no statically known variants")
		}
		if operation.Control == "" {
			return altExpr(branches...)
		}
		control := primitiveAtom(operation.Control, path, "union(control="+operation.Control+")")
		control.Site = operation.Site
		return concatExpr(atomExpr(control), altExpr(branches...))
	case "conditional", "switch", "type_switch":
		if operation.Kind != "type_switch" && hasDiscriminants(operation) {
			branches := discriminantBranches(operation)
			alternatives := make([]shapeExpr, 0, len(branches))
			for _, branch := range branches {
				alternatives = append(alternatives, concatExpr(atomExpr(branch.marker), sourceSequenceExpr(branch.ops)))
			}
			return altExpr(alternatives...)
		}
		branches := make([]shapeExpr, 0, len(operation.Variants)+1)
		for _, variant := range operation.Variants {
			branches = append(branches, sourceSequenceExpr(variant.Ops))
		}
		if operation.HasDefault {
			branches = append(branches, sourceSequenceExpr(operation.Default))
		}
		if len(branches) == 0 {
			return unknownExpr("gophertunnel: control-flow operation at " + path + " has no finite paths")
		}
		return altExpr(branches...)
	case "recursive":
		return recurseExpr(path)
	case "unresolved":
		reason := operation.Reason
		if reason == "" {
			reason = "operation is not statically resolvable"
		}
		where := path
		if operation.Site != "" {
			where += " (" + operation.Site + ")"
		}
		return unknownExpr("gophertunnel: " + reason + " at " + where)
	default:
		return unknownExpr("gophertunnel: unsupported operation " + operation.Kind + " at " + path)
	}
}
