package emitrust

import (
	"fmt"
	"strconv"
	"strings"

	"protocolgen/internal/manifest"
)

// The manifest describes wire shape, not hostile-input policy, so collection
// bounds come from the reader at decode time rather than from a source claim.
const maxFixedArray = 32

// codecEmitter walks the same manifest nodes as the type emitter and writes the
// matching Encode/Decode bodies. It must agree with rustType on every shape
// decision: a mismatch is a compile error in the generated crate.
type codecEmitter struct {
	g *generator
}

func (e *codecEmitter) nativeCall(node manifest.Node) (string, bool, error) {
	typ, matched, err := e.g.nativeRustType(node)
	if err != nil {
		return "", true, err
	}
	return typ, matched, nil
}

// minWireSize reports the smallest number of bytes a value of this shape can
// occupy. Collection decoding uses it to reject a declared count that the
// remaining input could not possibly cover, before reserving.
func minWireSize(node manifest.Node, visiting map[string]bool) int {
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return 1
		}
		switch node.Primitive.Code {
		case "uuid":
			return 16
		case "nbt_le":
			return 1
		case "bool", "i8", "u8":
			return 1
		}
		if strings.HasPrefix(node.Primitive.Code, "var_") || strings.HasPrefix(node.Primitive.Code, "zigzag_") {
			return 1
		}
		return int(node.Primitive.Width) / 8
	case manifest.KindString, manifest.KindBytes, manifest.KindArray, manifest.KindMap,
		manifest.KindOptional, manifest.KindBitset:
		return 1
	case manifest.KindFixedArray:
		if node.Element == nil {
			return 0
		}
		return int(node.Length) * minWireSize(*node.Element, visiting)
	case manifest.KindStruct:
		if node.TypeID != "" {
			if visiting[node.TypeID] {
				return 1
			}
			visiting[node.TypeID] = true
			defer delete(visiting, node.TypeID)
		}
		total := 0
		for _, field := range node.Fields {
			total += minWireSize(field.Encode, visiting)
		}
		return total
	case manifest.KindEnum:
		if node.Primitive == nil {
			return 1
		}
		return minWireSize(manifest.Node{Kind: manifest.KindPrimitive, Primitive: node.Primitive}, visiting)
	case manifest.KindUnion:
		best := -1
		for _, variant := range node.Variants {
			size := minWireSize(variant.Encode, visiting)
			if best < 0 || size < best {
				best = size
			}
		}
		if best < 0 {
			best = 0
		}
		control := 1
		if node.Control != nil {
			control = minWireSize(*node.Control, visiting)
		}
		return control + best
	case manifest.KindRecursive:
		return 1
	case manifest.KindVoid:
		return 0
	default:
		return 1
	}
}

func minSizeOf(node manifest.Node) int {
	size := minWireSize(node, map[string]bool{})
	if size < 1 {
		return 1
	}
	return size
}

// encode writes the statements that serialise expr, which already has the Rust
// type rustType assigned to this node.
func (e *codecEmitter) encode(b *strings.Builder, node manifest.Node, expr, indent string) error {
	if _, matched, err := e.nativeCall(node); err != nil {
		return err
	} else if matched {
		fmt.Fprintf(b, "%s%s.encode(writer);\n", indent, expr)
		return nil
	}

	switch node.Kind {
	case manifest.KindPrimitive, manifest.KindEnum:
		fmt.Fprintf(b, "%s%s.encode(writer);\n", indent, expr)
		return e.encodeNumberConstraints(b, node, expr, indent)

	case manifest.KindString:
		reference := rustStringReference(expr)
		if min, max, ok := stringBounds(node); ok {
			fmt.Fprintf(b, "%swire::encode_string_limits(writer, %s, %d, %d);\n", indent, reference, min, max)
		} else {
			fmt.Fprintf(b, "%s%s.encode(writer);\n", indent, expr)
		}
		if node.Constraints != nil && node.Constraints.Pattern != "" {
			fmt.Fprintf(b, "%swire::assert_pattern(%s, %q);\n", indent, reference, node.Constraints.Pattern)
		}
		return nil

	case manifest.KindBytes:
		if min, max, ok := stringBounds(node); ok {
			fmt.Fprintf(b, "%swire::encode_bytes_limits(writer, %s.as_ref(), %d, %d);\n", indent, expr, min, max)
		} else {
			fmt.Fprintf(b, "%s%s.encode(writer);\n", indent, expr)
		}
		return nil

	case manifest.KindStruct, manifest.KindUnion, manifest.KindRecursive:
		fmt.Fprintf(b, "%s%s.encode(writer);\n", indent, expr)
		return nil

	case manifest.KindVoid:
		return nil

	case manifest.KindBitset:
		fmt.Fprintf(b, "%swire::encode_bitset(writer, %s.0.as_slice(), %d);\n", indent, expr, node.Length)
		return nil

	case manifest.KindOptional:
		return e.encodeOptional(b, node, expr, indent)

	case manifest.KindFixedArray:
		if node.Element == nil {
			return fmt.Errorf("fixed array has no element")
		}
		fmt.Fprintf(b, "%sfor item in %s.iter() {\n", indent, expr)
		if err := e.encode(b, *node.Element, "item", indent+"    "); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil

	case manifest.KindArray:
		if node.Element == nil {
			return fmt.Errorf("array has no element")
		}
		helper, err := collectionHelper(node, "encode_collection")
		if err != nil {
			return err
		}
		if e.directlyEncodable(*node.Element) {
			if min, max, ok := arrayBounds(node); ok && helper == "encode_collection" {
				fmt.Fprintf(b, "%swire::encode_collection_limits(writer, %s.as_slice(), %d, %d);\n", indent, expr, min, max)
			} else {
				e.encodeLengthAssertion(b, node, expr, indent)
				fmt.Fprintf(b, "%swire::%s(writer, %s.as_slice());\n", indent, helper, expr)
			}
			return nil
		}
		e.encodeLengthAssertion(b, node, expr, indent)
		if err := e.writeCollectionPrefix(b, node, expr, indent); err != nil {
			return err
		}
		fmt.Fprintf(b, "%sfor item in %s.iter() {\n", indent, expr)
		if err := e.encode(b, *node.Element, "item", indent+"    "); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil

	case manifest.KindMap:
		if node.Key == nil || node.Value == nil {
			return fmt.Errorf("map has no key/value")
		}
		if e.directlyEncodable(*node.Key) && e.directlyEncodable(*node.Value) {
			if min, max, ok := mapBounds(node); ok {
				fmt.Fprintf(b, "%swire::encode_map_limits(writer, %s.as_slice(), %d, %d);\n", indent, expr, min, max)
			} else {
				fmt.Fprintf(b, "%swire::encode_map(writer, %s.as_slice());\n", indent, expr)
			}
			return nil
		}
		e.encodeLengthAssertion(b, node, expr, indent)
		fmt.Fprintf(b, "%swriter.write_var_u32(%s.len() as u32);\n", indent, expr)
		fmt.Fprintf(b, "%sfor (key, value) in %s.iter() {\n", indent, expr)
		if err := e.encode(b, *node.Key, "key", indent+"    "); err != nil {
			return err
		}
		if err := e.encode(b, *node.Value, "value", indent+"    "); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil

	default:
		return fmt.Errorf("unsupported node kind %q in encode", node.Kind)
	}
}

func rustStringReference(expr string) string {
	if strings.HasPrefix(expr, "self.") {
		return "&" + expr
	}
	return expr
}

func (e *codecEmitter) encodeNumberConstraints(b *strings.Builder, node manifest.Node, expr, indent string) error {
	if node.Constraints == nil || node.Constraints.Minimum == nil && node.Constraints.Maximum == nil {
		return nil
	}
	raw, err := rustNumberExpression(node, expr)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%swire::assert_number_limits(%s, %s, %s);\n", indent, raw, rustOptionNumber(node, node.Constraints.Minimum), rustOptionNumber(node, node.Constraints.Maximum))
	return nil
}

func (e *codecEmitter) encodeLengthAssertion(b *strings.Builder, node manifest.Node, expr, indent string) {
	var min, max uint64
	var ok bool
	if node.Kind == manifest.KindMap {
		min, max, ok = mapBounds(node)
	} else {
		min, max, ok = arrayBounds(node)
	}
	if ok {
		fmt.Fprintf(b, "%swire::assert_length(%s.len(), %d, %d);\n", indent, expr, min, max)
	}
}

func (e *codecEmitter) writeCollectionPrefix(b *strings.Builder, node manifest.Node, expr, indent string) error {
	code, err := prefixCode(node)
	if err != nil {
		return err
	}
	switch code {
	case "var_u32":
		fmt.Fprintf(b, "%swriter.write_var_u32(%s.len() as u32);\n", indent, expr)
	case "u32le":
		fmt.Fprintf(b, "%swriter.write_all(&(%s.len() as u32).to_le_bytes());\n", indent, expr)
	default:
		return fmt.Errorf("unsupported collection prefix %q", code)
	}
	return nil
}

// encodeOptional writes the presence marker(s). Cereal wraps some optional
// values in a second, always-present optional; both markers are on the wire
// even though the Rust type exposes only the inner state.
func (e *codecEmitter) encodeOptional(b *strings.Builder, node manifest.Node, expr, indent string) error {
	if node.Value == nil {
		return fmt.Errorf("optional has no value")
	}
	value := *node.Value
	nested := value.Kind == manifest.KindOptional
	if nested {
		if value.Value == nil {
			return fmt.Errorf("nested optional has no value")
		}
		value = *value.Value
		fmt.Fprintf(b, "%swriter.write_u8(1);\n", indent)
	}
	fmt.Fprintf(b, "%smatch &%s {\n", indent, expr)
	fmt.Fprintf(b, "%s    Some(value) => {\n", indent)
	fmt.Fprintf(b, "%s        writer.write_u8(1);\n", indent)
	if err := e.encode(b, value, "value", indent+"        "); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s    }\n", indent)
	fmt.Fprintf(b, "%s    None => writer.write_u8(0),\n", indent)
	fmt.Fprintf(b, "%s}\n", indent)
	return nil
}

// directlyEncodable reports whether a node's Rust type implements Encode and
// Decode on its own, so a generic collection helper can be used instead of an
// emitted loop.
func (e *codecEmitter) directlyEncodable(node manifest.Node) bool {
	if _, matched, err := e.nativeCall(node); err == nil && matched {
		return true
	}
	switch node.Kind {
	case manifest.KindPrimitive, manifest.KindString, manifest.KindBytes,
		manifest.KindStruct, manifest.KindEnum, manifest.KindUnion, manifest.KindRecursive:
		return true
	default:
		return false
	}
}

func prefixCode(node manifest.Node) (string, error) {
	if node.Prefix == nil || node.Prefix.Primitive == nil {
		return "", fmt.Errorf("%s node has no primitive prefix", node.Kind)
	}
	return node.Prefix.Primitive.Code, nil
}

func collectionHelper(node manifest.Node, base string) (string, error) {
	code, err := prefixCode(node)
	if err != nil {
		return "", err
	}
	switch code {
	case "var_u32":
		return base, nil
	case "u32le":
		return base + "_u32le", nil
	default:
		return "", fmt.Errorf("unsupported collection prefix %q", code)
	}
}

// decode returns a Rust expression that yields the decoded value. The
// expression may be a block, so it is always safe to use in value position.
func (e *codecEmitter) decode(node manifest.Node, hint, indent string) (string, error) {
	if typ, matched, err := e.nativeCall(node); err != nil {
		return "", err
	} else if matched {
		return decodeCall(typ), nil
	}

	switch node.Kind {
	case manifest.KindPrimitive, manifest.KindEnum:
		typ, err := e.g.rustType(node, hint)
		if err != nil {
			return "", err
		}
		decoded := decodeCall(typ)
		if node.Constraints == nil || node.Constraints.Minimum == nil && node.Constraints.Maximum == nil {
			return decoded, nil
		}
		raw, err := rustNumberExpression(node, "value")
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("{ let value = %s; wire::validate_number_limits(%s, %s, %s)?; value }", decoded, raw, rustOptionNumber(node, node.Constraints.Minimum), rustOptionNumber(node, node.Constraints.Maximum)), nil

	case manifest.KindString:
		decoded := decodeCall("String")
		if min, max, ok := stringBounds(node); ok {
			decoded = fmt.Sprintf("wire::decode_string_limits(reader, %d, %d)?", min, max)
		}
		if node.Constraints != nil && node.Constraints.Pattern != "" {
			return fmt.Sprintf("{ let value = %s; wire::validate_pattern(&value, %q)?; value }", decoded, node.Constraints.Pattern), nil
		}
		return decoded, nil

	case manifest.KindBytes:
		if min, max, ok := stringBounds(node); ok {
			return fmt.Sprintf("wire::decode_bytes_limits(reader, %d, %d)?", min, max), nil
		}
		typ, err := e.g.rustType(node, hint)
		if err != nil {
			return "", err
		}
		return decodeCall(typ), nil

	case manifest.KindStruct, manifest.KindUnion, manifest.KindRecursive:
		typ, err := e.g.rustType(node, hint)
		if err != nil {
			return "", err
		}
		return decodeCall(typ), nil

	case manifest.KindVoid:
		return "()", nil

	case manifest.KindBitset:
		words := (node.Length + 63) / 64
		return fmt.Sprintf("%s(wire::decode_bitset::<%d>(reader, %d)?)", e.bitsetName(node), words, node.Length), nil

	case manifest.KindOptional:
		return e.decodeOptional(node, hint, indent)

	case manifest.KindFixedArray:
		if node.Element == nil {
			return "", fmt.Errorf("fixed array has no element")
		}
		if node.Length > maxFixedArray {
			return "", fmt.Errorf("fixed array of %d exceeds the emitted literal bound", node.Length)
		}
		element, err := e.decode(*node.Element, hint+"Item", indent)
		if err != nil {
			return "", err
		}
		parts := make([]string, 0, node.Length)
		for i := uint64(0); i < node.Length; i++ {
			parts = append(parts, element)
		}
		// Rust evaluates array literal elements left to right, which is the
		// wire order.
		return "[" + strings.Join(parts, ", ") + "]", nil

	case manifest.KindArray:
		if node.Element == nil {
			return "", fmt.Errorf("array has no element")
		}
		helper, err := collectionHelper(node, "decode_collection")
		if err != nil {
			return "", err
		}
		min := minSizeOf(*node.Element)
		element, err := e.g.rustType(*node.Element, hint+"Item")
		if err != nil {
			return "", err
		}
		if e.directlyEncodable(*node.Element) {
			if lower, upper, ok := arrayBounds(node); ok && helper == "decode_collection" {
				return fmt.Sprintf("wire::decode_collection_limits::<%s>(reader, %d, %d, %d)?", element, min, lower, upper), nil
			}
			return fmt.Sprintf("wire::%s::<%s>(reader, %d)?", helper, element, min), nil
		}
		return e.decodeCollectionLoop(node, helper, element, min, hint, indent)

	case manifest.KindMap:
		if node.Key == nil || node.Value == nil {
			return "", fmt.Errorf("map has no key/value")
		}
		min := minSizeOf(*node.Key) + minSizeOf(*node.Value)
		keyType, err := e.g.rustType(*node.Key, hint+"Key")
		if err != nil {
			return "", err
		}
		valueType, err := e.g.rustType(*node.Value, hint+"Value")
		if err != nil {
			return "", err
		}
		if e.directlyEncodable(*node.Key) && e.directlyEncodable(*node.Value) {
			if lower, upper, ok := mapBounds(node); ok {
				return fmt.Sprintf("wire::decode_map_limits::<%s, %s>(reader, %d, %d, %d)?", keyType, valueType, min, lower, upper), nil
			}
			return fmt.Sprintf("wire::decode_map::<%s, %s>(reader, %d)?", keyType, valueType, min), nil
		}
		key, err := e.decode(*node.Key, hint+"Key", indent+"        ")
		if err != nil {
			return "", err
		}
		value, err := e.decode(*node.Value, hint+"Value", indent+"        ")
		if err != nil {
			return "", err
		}
		var b strings.Builder
		fmt.Fprintf(&b, "{\n")
		fmt.Fprintf(&b, "%s    let declared = u64::from(reader.read_var_u32()?);\n", indent)
		if lower, upper, ok := mapBounds(node); ok {
			fmt.Fprintf(&b, "%s    let count = reader.checked_count_limits(declared, %d, %d, %d)?;\n", indent, min, lower, upper)
		} else {
			fmt.Fprintf(&b, "%s    let count = reader.checked_count(declared, %d)?;\n", indent, min)
		}
		fmt.Fprintf(&b, "%s    let mut out: Vec<(%s, %s)> = Vec::with_capacity(count);\n", indent, keyType, valueType)
		fmt.Fprintf(&b, "%s    for _ in 0..count {\n", indent)
		fmt.Fprintf(&b, "%s        let key = %s;\n", indent, key)
		fmt.Fprintf(&b, "%s        let value = %s;\n", indent, value)
		fmt.Fprintf(&b, "%s        out.push((key, value));\n", indent)
		fmt.Fprintf(&b, "%s    }\n", indent)
		fmt.Fprintf(&b, "%s    out\n", indent)
		fmt.Fprintf(&b, "%s}", indent)
		return b.String(), nil

	default:
		return "", fmt.Errorf("unsupported node kind %q in decode", node.Kind)
	}
}

func (e *codecEmitter) decodeCollectionLoop(node manifest.Node, helper, element string, min int, hint, indent string) (string, error) {
	item, err := e.decode(*node.Element, hint+"Item", indent+"        ")
	if err != nil {
		return "", err
	}
	code, err := prefixCode(node)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	fmt.Fprintf(&b, "{\n")
	switch code {
	case "var_u32":
		fmt.Fprintf(&b, "%s    let declared = u64::from(reader.read_var_u32()?);\n", indent)
	case "u32le":
		fmt.Fprintf(&b, "%s    let declared = u64::from(u32::from_le_bytes(reader.read_bytes::<4>()?));\n", indent)
	default:
		return "", fmt.Errorf("unsupported collection prefix %q", code)
	}
	if lower, upper, ok := arrayBounds(node); ok {
		fmt.Fprintf(&b, "%s    let count = reader.checked_count_limits(declared, %d, %d, %d)?;\n", indent, min, lower, upper)
	} else {
		fmt.Fprintf(&b, "%s    let count = reader.checked_count(declared, %d)?;\n", indent, min)
	}
	fmt.Fprintf(&b, "%s    let mut out: Vec<%s> = Vec::with_capacity(count);\n", indent, element)
	fmt.Fprintf(&b, "%s    for _ in 0..count {\n", indent)
	fmt.Fprintf(&b, "%s        out.push(%s);\n", indent, item)
	fmt.Fprintf(&b, "%s    }\n", indent)
	fmt.Fprintf(&b, "%s    out\n", indent)
	fmt.Fprintf(&b, "%s}", indent)
	return b.String(), nil
}

func rustNumberExpression(node manifest.Node, expr string) (string, error) {
	if node.Kind == manifest.KindEnum {
		return expr + ".to_raw()", nil
	}
	if node.Kind != manifest.KindPrimitive || node.Primitive == nil {
		return "", fmt.Errorf("numeric constraints require primitive or enum")
	}
	if node.Primitive.Code == "bool" || node.Primitive.Code == "uuid" || node.Primitive.Code == "nbt_le" {
		return "", fmt.Errorf("primitive %s cannot carry numeric constraints", node.Primitive.Code)
	}
	return expr + ".0", nil
}

func rustOptionNumber(node manifest.Node, value *float64) string {
	if value == nil {
		return "None"
	}
	literal := strconv.FormatFloat(*value, 'f', -1, 64)
	if node.Primitive != nil {
		switch node.Primitive.Code {
		case "f32le", "f32be", "f64le", "f64be":
			if !strings.Contains(literal, ".") {
				literal += ".0"
			}
		}
	}
	return "Some(" + literal + ")"
}

func stringBounds(node manifest.Node) (uint64, uint64, bool) {
	if node.Constraints == nil || node.Constraints.MinLength == nil && node.Constraints.MaxLength == nil {
		return 0, 0, false
	}
	min, max := uint64(0), ^uint64(0)
	if node.Constraints.MinLength != nil {
		min = *node.Constraints.MinLength
	}
	if node.Constraints.MaxLength != nil {
		max = *node.Constraints.MaxLength
	}
	return min, max, true
}

func arrayBounds(node manifest.Node) (uint64, uint64, bool) {
	if node.Constraints == nil || node.Constraints.MinItems == nil && node.Constraints.MaxItems == nil {
		return 0, 0, false
	}
	min, max := uint64(0), ^uint64(0)
	if node.Constraints.MinItems != nil {
		min = *node.Constraints.MinItems
	}
	if node.Constraints.MaxItems != nil {
		max = *node.Constraints.MaxItems
	}
	return min, max, true
}

func mapBounds(node manifest.Node) (uint64, uint64, bool) {
	if node.Constraints == nil || node.Constraints.MinProperties == nil && node.Constraints.MaxProperties == nil {
		return 0, 0, false
	}
	min, max := uint64(0), ^uint64(0)
	if node.Constraints.MinProperties != nil {
		min = *node.Constraints.MinProperties
	}
	if node.Constraints.MaxProperties != nil {
		max = *node.Constraints.MaxProperties
	}
	return min, max, true
}

func (e *codecEmitter) decodeOptional(node manifest.Node, hint, indent string) (string, error) {
	if node.Value == nil {
		return "", fmt.Errorf("optional has no value")
	}
	value := *node.Value
	nested := value.Kind == manifest.KindOptional
	if nested {
		if value.Value == nil {
			return "", fmt.Errorf("nested optional has no value")
		}
		value = *value.Value
	}
	inner, err := e.decode(value, hint+"Value", indent+"        ")
	if err != nil {
		return "", err
	}
	var b strings.Builder
	fmt.Fprintf(&b, "{\n")
	if nested {
		fmt.Fprintf(&b, "%s    if reader.read_u8()? != 0 && reader.read_u8()? != 0 {\n", indent)
		fmt.Fprintf(&b, "%s        Some(%s)\n", indent, inner)
		fmt.Fprintf(&b, "%s    } else {\n", indent)
		fmt.Fprintf(&b, "%s        None\n", indent)
		fmt.Fprintf(&b, "%s    }\n", indent)
	} else {
		fmt.Fprintf(&b, "%s    if reader.read_u8()? == 0 {\n", indent)
		fmt.Fprintf(&b, "%s        None\n", indent)
		fmt.Fprintf(&b, "%s    } else {\n", indent)
		fmt.Fprintf(&b, "%s        Some(%s)\n", indent, inner)
		fmt.Fprintf(&b, "%s    }\n", indent)
	}
	fmt.Fprintf(&b, "%s}", indent)
	return b.String(), nil
}

// decodeCall names the type at the call site so a disagreement between the
// type emitter and this walk is a compile error rather than silent inference.
func decodeCall(typ string) string {
	return fmt.Sprintf("<%s as wire::Decode>::decode(reader)?", typ)
}

func (e *codecEmitter) bitsetName(node manifest.Node) string {
	return fmt.Sprintf("Bitset%d", node.Length)
}
