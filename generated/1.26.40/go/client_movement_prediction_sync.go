// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientMovementPredictionSync struct {
	ActorDataFlag      ActorDataFlagComponent
	ActorBoundingBox   ActorDataBoundingBoxComponent
	MovementAttributes [9]float32
	ActorUniqueID      ActorUniqueID
	ActorFlyingState   bool
}

func (p *ClientMovementPredictionSync) Encode(w Encoder) error {
	if err := w.Write("ClientMovementPredictionSyncPacket.Actor Data Flag", Shape{Kind: "struct", Semantic: "ActorDataFlagComponent", TypeID: "ActorDataFlagComponent", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Flag Bitset Data", Shape: Shape{Kind: "bitset", Representation: "bitset", Length: 131}}}}, p.ActorDataFlag); err != nil {
		return err
	}
	if err := w.Write("ClientMovementPredictionSyncPacket.Actor Bounding Box", Shape{Kind: "struct", Semantic: "ActorDataBoundingBoxComponent", TypeID: "ActorDataBoundingBoxComponent", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Data Bounding Box", Shape: Shape{Kind: "fixed_array", Length: 3, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.ActorBoundingBox); err != nil {
		return err
	}
	if err := w.Write("ClientMovementPredictionSyncPacket.Movement Attributes", Shape{Kind: "fixed_array", Semantic: "MovementAttributesComponent", TypeID: "MovementAttributesComponent.json#", Length: 9, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, p.MovementAttributes); err != nil {
		return err
	}
	if err := w.Write("ClientMovementPredictionSyncPacket.Actor Unique ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.ActorUniqueID); err != nil {
		return err
	}
	if err := w.Write("ClientMovementPredictionSyncPacket.Actor Flying State", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ActorFlyingState); err != nil {
		return err
	}
	return nil
}

func DecodeClientMovementPredictionSync(r Decoder) (ClientMovementPredictionSync, error) {
	var p ClientMovementPredictionSync
	{
		raw, err := r.Read("ClientMovementPredictionSyncPacket.Actor Data Flag", Shape{Kind: "struct", Semantic: "ActorDataFlagComponent", TypeID: "ActorDataFlagComponent", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Flag Bitset Data", Shape: Shape{Kind: "bitset", Representation: "bitset", Length: 131}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorDataFlagComponent)
		if !ok {
			return p, fmt.Errorf("field ClientMovementPredictionSyncPacket.Actor Data Flag has unexpected decoded type %T", raw)
		}
		p.ActorDataFlag = value
	}
	{
		raw, err := r.Read("ClientMovementPredictionSyncPacket.Actor Bounding Box", Shape{Kind: "struct", Semantic: "ActorDataBoundingBoxComponent", TypeID: "ActorDataBoundingBoxComponent", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Data Bounding Box", Shape: Shape{Kind: "fixed_array", Length: 3, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorDataBoundingBoxComponent)
		if !ok {
			return p, fmt.Errorf("field ClientMovementPredictionSyncPacket.Actor Bounding Box has unexpected decoded type %T", raw)
		}
		p.ActorBoundingBox = value
	}
	{
		raw, err := r.Read("ClientMovementPredictionSyncPacket.Movement Attributes", Shape{Kind: "fixed_array", Semantic: "MovementAttributesComponent", TypeID: "MovementAttributesComponent.json#", Length: 9, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([9]float32)
		if !ok {
			return p, fmt.Errorf("field ClientMovementPredictionSyncPacket.Movement Attributes has unexpected decoded type %T", raw)
		}
		p.MovementAttributes = value
	}
	{
		raw, err := r.Read("ClientMovementPredictionSyncPacket.Actor Unique ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field ClientMovementPredictionSyncPacket.Actor Unique ID has unexpected decoded type %T", raw)
		}
		p.ActorUniqueID = value
	}
	{
		raw, err := r.Read("ClientMovementPredictionSyncPacket.Actor Flying State", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ClientMovementPredictionSyncPacket.Actor Flying State has unexpected decoded type %T", raw)
		}
		p.ActorFlyingState = value
	}
	return p, nil
}
