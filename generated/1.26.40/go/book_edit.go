// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type BookEdit struct {
	BookSlot  int32
	Operation BookEditAction
}

func (p *BookEdit) Encode(w Encoder) error {
	if err := w.Write("BookEditPacket.Book Slot", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.BookSlot); err != nil {
		return err
	}
	if err := w.Write("BookEditPacket.Operation", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "BookEditAction::ReplacePage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::ReplacePage", TypeID: "BookEditAction::ReplacePage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Page Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Photo Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "BookEditAction::AddPage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::AddPage", TypeID: "BookEditAction::AddPage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Page Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Photo Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "BookEditAction::DeletePage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::DeletePage", TypeID: "BookEditAction::DeletePage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "BookEditAction::SwapPages", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::SwapPages", TypeID: "BookEditAction::SwapPages", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Swap With Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 4, Name: "BookEditAction::Finalize", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::Finalize", TypeID: "BookEditAction::Finalize", Fields: []ShapeField{{Ordinal: 0, Name: "Title", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Author", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "XUID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.Operation); err != nil {
		return err
	}
	return nil
}

func DecodeBookEdit(r Decoder) (BookEdit, error) {
	var p BookEdit
	{
		raw, err := r.Read("BookEditPacket.Book Slot", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field BookEditPacket.Book Slot has unexpected decoded type %T", raw)
		}
		p.BookSlot = value
	}
	{
		raw, err := r.Read("BookEditPacket.Operation", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "BookEditAction::ReplacePage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::ReplacePage", TypeID: "BookEditAction::ReplacePage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Page Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Photo Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "BookEditAction::AddPage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::AddPage", TypeID: "BookEditAction::AddPage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Page Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Photo Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "BookEditAction::DeletePage", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::DeletePage", TypeID: "BookEditAction::DeletePage", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "BookEditAction::SwapPages", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::SwapPages", TypeID: "BookEditAction::SwapPages", Fields: []ShapeField{{Ordinal: 0, Name: "Page Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Swap With Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 4, Name: "BookEditAction::Finalize", Shape: Shape{Kind: "struct", Semantic: "BookEditAction::Finalize", TypeID: "BookEditAction::Finalize", Fields: []ShapeField{{Ordinal: 0, Name: "Title", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Author", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "XUID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BookEditAction)
		if !ok {
			return p, fmt.Errorf("field BookEditPacket.Operation has unexpected decoded type %T", raw)
		}
		p.Operation = value
	}
	return p, nil
}
