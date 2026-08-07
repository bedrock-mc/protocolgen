// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RemoveVolumeEntity struct {
	EntityNetworkId EntityNetId
	DimensionType   DimensionType
}

func (p *RemoveVolumeEntity) Encode(w Encoder) error {
	if err := w.Write("RemoveVolumeEntityPacket.Entity Network Id", Shape{Kind: "struct", Semantic: "EntityNetId", TypeID: "EntityNetId", Fields: []ShapeField{{Ordinal: 0, Name: "Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, p.EntityNetworkId); err != nil {
		return err
	}
	if err := w.Write("RemoveVolumeEntityPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionType); err != nil {
		return err
	}
	return nil
}

func DecodeRemoveVolumeEntity(r Decoder) (RemoveVolumeEntity, error) {
	var p RemoveVolumeEntity
	{
		raw, err := r.Read("RemoveVolumeEntityPacket.Entity Network Id", Shape{Kind: "struct", Semantic: "EntityNetId", TypeID: "EntityNetId", Fields: []ShapeField{{Ordinal: 0, Name: "Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(EntityNetId)
		if !ok {
			return p, fmt.Errorf("field RemoveVolumeEntityPacket.Entity Network Id has unexpected decoded type %T", raw)
		}
		p.EntityNetworkId = value
	}
	{
		raw, err := r.Read("RemoveVolumeEntityPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field RemoveVolumeEntityPacket.Dimension Type has unexpected decoded type %T", raw)
		}
		p.DimensionType = value
	}
	return p, nil
}
