// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CommandBlockUpdate struct {
	Target             CommandBlockUpdateTarget
	Command            string
	LastOutput         string
	Name               string
	FilteredName       string
	TrackOutput        bool
	TickDelay          int32
	ExecuteOnFirstTick bool
}

func (p *CommandBlockUpdate) Encode(w Encoder) error {
	if err := w.Write("CommandBlockUpdatePacket.Target", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "CommandBlockUpdatePacketPayload::EntityCommandTarget", Shape: Shape{Kind: "struct", Semantic: "CommandBlockUpdatePacketPayload::EntityCommandTarget", TypeID: "CommandBlockUpdatePacketPayload::EntityCommandTarget", Fields: []ShapeField{{Ordinal: 0, Name: "Target Runtime ID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}}}}, {Value: 1, Name: "CommandBlockUpdatePacketPayload::BlockCommandData", Shape: Shape{Kind: "struct", Semantic: "CommandBlockUpdatePacketPayload::BlockCommandData", TypeID: "CommandBlockUpdatePacketPayload::BlockCommandData", Fields: []ShapeField{{Ordinal: 0, Name: "Block Position", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Command Block Mode", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Redstone Mode", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Is Conditional", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}}, p.Target); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Command); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Last Output", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.LastOutput); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Name); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.FilteredName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FilteredName); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Track Output", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.TrackOutput); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Tick Delay", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.TickDelay); err != nil {
		return err
	}
	if err := w.Write("CommandBlockUpdatePacket.Execute On First Tick", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ExecuteOnFirstTick); err != nil {
		return err
	}
	return nil
}

func DecodeCommandBlockUpdate(r Decoder) (CommandBlockUpdate, error) {
	var p CommandBlockUpdate
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Target", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "CommandBlockUpdatePacketPayload::EntityCommandTarget", Shape: Shape{Kind: "struct", Semantic: "CommandBlockUpdatePacketPayload::EntityCommandTarget", TypeID: "CommandBlockUpdatePacketPayload::EntityCommandTarget", Fields: []ShapeField{{Ordinal: 0, Name: "Target Runtime ID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}}}}, {Value: 1, Name: "CommandBlockUpdatePacketPayload::BlockCommandData", Shape: Shape{Kind: "struct", Semantic: "CommandBlockUpdatePacketPayload::BlockCommandData", TypeID: "CommandBlockUpdatePacketPayload::BlockCommandData", Fields: []ShapeField{{Ordinal: 0, Name: "Block Position", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Command Block Mode", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Redstone Mode", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Is Conditional", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CommandBlockUpdateTarget)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Target has unexpected decoded type %T", raw)
		}
		p.Target = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Command has unexpected decoded type %T", raw)
		}
		p.Command = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Last Output", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Last Output has unexpected decoded type %T", raw)
		}
		p.LastOutput = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Name has unexpected decoded type %T", raw)
		}
		p.Name = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.FilteredName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.FilteredName has unexpected decoded type %T", raw)
		}
		p.FilteredName = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Track Output", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Track Output has unexpected decoded type %T", raw)
		}
		p.TrackOutput = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Tick Delay", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Tick Delay has unexpected decoded type %T", raw)
		}
		p.TickDelay = value
	}
	{
		raw, err := r.Read("CommandBlockUpdatePacket.Execute On First Tick", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CommandBlockUpdatePacket.Execute On First Tick has unexpected decoded type %T", raw)
		}
		p.ExecuteOnFirstTick = value
	}
	return p, nil
}
