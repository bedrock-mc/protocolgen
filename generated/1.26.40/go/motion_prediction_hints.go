// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type MotionPredictionHints struct {
	MRuntimeId ActorRuntimeID
	MMotion    mgl32.Vec3
	MOnGround  bool
}

func (p *MotionPredictionHints) Encode(w Encoder) error {
	if err := w.Write("MotionPredictionHintsPacket.mRuntimeId", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.MRuntimeId); err != nil {
		return err
	}
	if err := w.Write("MotionPredictionHintsPacket.mMotion", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.MMotion); err != nil {
		return err
	}
	if err := w.Write("MotionPredictionHintsPacket.mOnGround", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.MOnGround); err != nil {
		return err
	}
	return nil
}

func DecodeMotionPredictionHints(r Decoder) (MotionPredictionHints, error) {
	var p MotionPredictionHints
	{
		raw, err := r.Read("MotionPredictionHintsPacket.mRuntimeId", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MotionPredictionHintsPacket.mRuntimeId has unexpected decoded type %T", raw)
		}
		p.MRuntimeId = value
	}
	{
		raw, err := r.Read("MotionPredictionHintsPacket.mMotion", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field MotionPredictionHintsPacket.mMotion has unexpected decoded type %T", raw)
		}
		p.MMotion = value
	}
	{
		raw, err := r.Read("MotionPredictionHintsPacket.mOnGround", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field MotionPredictionHintsPacket.mOnGround has unexpected decoded type %T", raw)
		}
		p.MOnGround = value
	}
	return p, nil
}
