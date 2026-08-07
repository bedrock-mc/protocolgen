// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AddVolumeEntity struct {
	EntityNetworkId EntityNetId
	Components      []byte
	JSONIdentifier  string
	InstanceName    string
	MinBounds       BlockPos
	MaxBounds       BlockPos
	DimensionType   DimensionType
	EngineVersion   string
}

func (p *AddVolumeEntity) Encode(w Encoder) error {
	if err := w.Write("AddVolumeEntityPacket.Entity Network Id", Shape{Kind: "struct", Semantic: "EntityNetId", TypeID: "EntityNetId", Fields: []ShapeField{{Ordinal: 0, Name: "Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, p.EntityNetworkId); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Components", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.Components); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.JSON Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.JSONIdentifier); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Instance Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.InstanceName); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Min Bounds", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.MinBounds); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Max Bounds", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.MaxBounds); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionType); err != nil {
		return err
	}
	if err := w.Write("AddVolumeEntityPacket.Engine Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.EngineVersion); err != nil {
		return err
	}
	return nil
}

func DecodeAddVolumeEntity(r Decoder) (AddVolumeEntity, error) {
	var p AddVolumeEntity
	{
		raw, err := r.Read("AddVolumeEntityPacket.Entity Network Id", Shape{Kind: "struct", Semantic: "EntityNetId", TypeID: "EntityNetId", Fields: []ShapeField{{Ordinal: 0, Name: "Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(EntityNetId)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Entity Network Id has unexpected decoded type %T", raw)
		}
		p.EntityNetworkId = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Components", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Components has unexpected decoded type %T", raw)
		}
		p.Components = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.JSON Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.JSON Identifier has unexpected decoded type %T", raw)
		}
		p.JSONIdentifier = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Instance Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Instance Name has unexpected decoded type %T", raw)
		}
		p.InstanceName = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Min Bounds", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Min Bounds has unexpected decoded type %T", raw)
		}
		p.MinBounds = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Max Bounds", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Max Bounds has unexpected decoded type %T", raw)
		}
		p.MaxBounds = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Dimension Type has unexpected decoded type %T", raw)
		}
		p.DimensionType = value
	}
	{
		raw, err := r.Read("AddVolumeEntityPacket.Engine Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddVolumeEntityPacket.Engine Version has unexpected decoded type %T", raw)
		}
		p.EngineVersion = value
	}
	return p, nil
}
