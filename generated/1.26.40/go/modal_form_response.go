// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ModalFormResponse struct {
	FormID           uint32
	JSONResponse     *string
	FormCancelReason *ModalFormCancelReason
}

func (p *ModalFormResponse) Encode(w Encoder) error {
	if err := w.Write("ModalFormResponsePacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.FormID); err != nil {
		return err
	}
	if err := w.Write("ModalFormResponsePacket.JSON Response", Shape{Kind: "optional", Value: &Shape{Kind: "string", Semantic: "Json::Value", TypeID: "Json__Value.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.JSONResponse); err != nil {
		return err
	}
	if err := w.Write("ModalFormResponsePacket.Form Cancel Reason", Shape{Kind: "optional", Value: &Shape{Kind: "enum", Semantic: "ModalFormCancelReason", TypeID: "enums/ModalFormCancelReason", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "UserClosed", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "UserBusy", Shape: Shape{Kind: "void"}}}}}, p.FormCancelReason); err != nil {
		return err
	}
	return nil
}

func DecodeModalFormResponse(r Decoder) (ModalFormResponse, error) {
	var p ModalFormResponse
	{
		raw, err := r.Read("ModalFormResponsePacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ModalFormResponsePacket.Form ID has unexpected decoded type %T", raw)
		}
		p.FormID = value
	}
	{
		raw, err := r.Read("ModalFormResponsePacket.JSON Response", Shape{Kind: "optional", Value: &Shape{Kind: "string", Semantic: "Json::Value", TypeID: "Json__Value.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*string)
		if !ok {
			return p, fmt.Errorf("field ModalFormResponsePacket.JSON Response has unexpected decoded type %T", raw)
		}
		p.JSONResponse = value
	}
	{
		raw, err := r.Read("ModalFormResponsePacket.Form Cancel Reason", Shape{Kind: "optional", Value: &Shape{Kind: "enum", Semantic: "ModalFormCancelReason", TypeID: "enums/ModalFormCancelReason", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "UserClosed", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "UserBusy", Shape: Shape{Kind: "void"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ModalFormCancelReason)
		if !ok {
			return p, fmt.Errorf("field ModalFormResponsePacket.Form Cancel Reason has unexpected decoded type %T", raw)
		}
		p.FormCancelReason = value
	}
	return p, nil
}
