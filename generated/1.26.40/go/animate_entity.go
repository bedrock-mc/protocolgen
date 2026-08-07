// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AnimateEntity struct {
	MAnimation             string
	MNextState             string
	MStopExpression        string
	MStopExpressionVersion int32
	MController            string
	MBlendOutTime          float32
	MRuntimeIds            []ActorRuntimeID
}

func (p *AnimateEntity) Encode(w Encoder) error {
	if err := w.Write("AnimateEntityPacket.mAnimation", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MAnimation); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mNextState", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MNextState); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mStopExpression", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MStopExpression); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mStopExpressionVersion", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.MStopExpressionVersion); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mController", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MController); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mBlendOutTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.MBlendOutTime); err != nil {
		return err
	}
	if err := w.Write("AnimateEntityPacket.mRuntimeIds", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}, p.MRuntimeIds); err != nil {
		return err
	}
	return nil
}

func DecodeAnimateEntity(r Decoder) (AnimateEntity, error) {
	var p AnimateEntity
	{
		raw, err := r.Read("AnimateEntityPacket.mAnimation", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mAnimation has unexpected decoded type %T", raw)
		}
		p.MAnimation = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mNextState", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mNextState has unexpected decoded type %T", raw)
		}
		p.MNextState = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mStopExpression", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mStopExpression has unexpected decoded type %T", raw)
		}
		p.MStopExpression = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mStopExpressionVersion", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mStopExpressionVersion has unexpected decoded type %T", raw)
		}
		p.MStopExpressionVersion = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mController", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mController has unexpected decoded type %T", raw)
		}
		p.MController = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mBlendOutTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mBlendOutTime has unexpected decoded type %T", raw)
		}
		p.MBlendOutTime = value
	}
	{
		raw, err := r.Read("AnimateEntityPacket.mRuntimeIds", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field AnimateEntityPacket.mRuntimeIds has unexpected decoded type %T", raw)
		}
		p.MRuntimeIds = value
	}
	return p, nil
}
