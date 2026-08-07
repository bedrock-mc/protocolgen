// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundMapItemData struct {
	MapID           ActorUniqueID
	Dimension       uint8
	IsLocked        bool
	MapOrigin       BlockPos
	CreationMapIDs  *[]ActorUniqueID
	Scale           *int8
	TrackedActorIDs *[]MapItemTrackedActorUniqueId
	Decorations     *[]MapDecoration
	Width           *int32
	Height          *int32
	StartX          *int32
	StartY          *int32
	Pixels          *[]uint32
}

func (p *ClientboundMapItemData) Encode(w Encoder) error {
	if err := w.Write("ClientboundMapItemDataPacket.Map ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.MapID); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Dimension", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.Dimension); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Is Locked", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsLocked); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Map Origin", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.MapOrigin); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Creation Map IDs", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}}, p.CreationMapIDs); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Scale", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}, p.Scale); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Tracked Actor IDs", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MapItemTrackedActor::UniqueId", TypeID: "MapItemTrackedActor::UniqueId", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "MapItemTrackedActor::Type", TypeID: "enums/MapItemTrackedActor::Type", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: 0, Name: "Entity", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "BlockEntity", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Other", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Entity ID", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}}, {Ordinal: 2, Name: "Block Position", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}}}, p.TrackedActorIDs); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Decorations", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MapDecoration", TypeID: "MapDecoration", Fields: []ShapeField{{Ordinal: 0, Name: "Image Type", Shape: Shape{Kind: "enum", Semantic: "MapDecoration::Type", TypeID: "enums/MapDecoration::Type", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "MarkerWhite", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "MarkerGreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "MarkerRed", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "MarkerBlue", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "XWhite", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "TriangleRed", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "SquareWhite", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "MarkerSign", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MarkerPink", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "MarkerOrange", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "MarkerYellow", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "MarkerTeal", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "TriangleGreen", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "SmallSquareWhite", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Mansion", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "Monument", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "NoDraw", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "VillageDesert", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "VillagePlains", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "VillageSavanna", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "VillageSnowy", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "VillageTaiga", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "JungleTemple", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "WitchHut", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "TrialChambers", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "Count", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Rotation", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 3, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 4, Name: "Label", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "Color", Shape: Shape{Kind: "struct", Semantic: "mce::Color", TypeID: "mce::Color", Fields: []ShapeField{{Ordinal: 0, Name: "Color", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}}}}}, p.Decorations); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Width", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, p.Width); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Height", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, p.Height); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Start X", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, p.StartX); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Start Y", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, p.StartY); err != nil {
		return err
	}
	if err := w.Write("ClientboundMapItemDataPacket.Pixels", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}, p.Pixels); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundMapItemData(r Decoder) (ClientboundMapItemData, error) {
	var p ClientboundMapItemData
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Map ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Map ID has unexpected decoded type %T", raw)
		}
		p.MapID = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Dimension", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Dimension has unexpected decoded type %T", raw)
		}
		p.Dimension = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Is Locked", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Is Locked has unexpected decoded type %T", raw)
		}
		p.IsLocked = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Map Origin", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Map Origin has unexpected decoded type %T", raw)
		}
		p.MapOrigin = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Creation Map IDs", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*[]ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Creation Map IDs has unexpected decoded type %T", raw)
		}
		p.CreationMapIDs = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Scale", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int8)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Scale has unexpected decoded type %T", raw)
		}
		p.Scale = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Tracked Actor IDs", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MapItemTrackedActor::UniqueId", TypeID: "MapItemTrackedActor::UniqueId", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "MapItemTrackedActor::Type", TypeID: "enums/MapItemTrackedActor::Type", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: 0, Name: "Entity", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "BlockEntity", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Other", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Entity ID", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}}, {Ordinal: 2, Name: "Block Position", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*[]MapItemTrackedActorUniqueId)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Tracked Actor IDs has unexpected decoded type %T", raw)
		}
		p.TrackedActorIDs = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Decorations", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MapDecoration", TypeID: "MapDecoration", Fields: []ShapeField{{Ordinal: 0, Name: "Image Type", Shape: Shape{Kind: "enum", Semantic: "MapDecoration::Type", TypeID: "enums/MapDecoration::Type", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "MarkerWhite", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "MarkerGreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "MarkerRed", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "MarkerBlue", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "XWhite", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "TriangleRed", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "SquareWhite", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "MarkerSign", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MarkerPink", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "MarkerOrange", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "MarkerYellow", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "MarkerTeal", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "TriangleGreen", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "SmallSquareWhite", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Mansion", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "Monument", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "NoDraw", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "VillageDesert", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "VillagePlains", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "VillageSavanna", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "VillageSnowy", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "VillageTaiga", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "JungleTemple", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "WitchHut", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "TrialChambers", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "Count", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Rotation", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 3, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 4, Name: "Label", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "Color", Shape: Shape{Kind: "struct", Semantic: "mce::Color", TypeID: "mce::Color", Fields: []ShapeField{{Ordinal: 0, Name: "Color", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*[]MapDecoration)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Decorations has unexpected decoded type %T", raw)
		}
		p.Decorations = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Width", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int32)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Width has unexpected decoded type %T", raw)
		}
		p.Width = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Height", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int32)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Height has unexpected decoded type %T", raw)
		}
		p.Height = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Start X", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int32)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Start X has unexpected decoded type %T", raw)
		}
		p.StartX = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Start Y", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int32)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Start Y has unexpected decoded type %T", raw)
		}
		p.StartY = value
	}
	{
		raw, err := r.Read("ClientboundMapItemDataPacket.Pixels", Shape{Kind: "optional", Value: &Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*[]uint32)
		if !ok {
			return p, fmt.Errorf("field ClientboundMapItemDataPacket.Pixels has unexpected decoded type %T", raw)
		}
		p.Pixels = value
	}
	return p, nil
}
