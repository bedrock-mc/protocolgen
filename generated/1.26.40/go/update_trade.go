// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateTrade struct {
	ContainerId       uint8
	Type              uint8
	Size              int32
	TraderTier        int32
	EntityUniqueId    ActorUniqueID
	LastTradingPlayer ActorUniqueID
	DisplayName       string
	UseNewTradeScreen bool
	UsingEconomyTrade bool
	Data              []byte
}

func (p *UpdateTrade) Encode(w Encoder) error {
	if err := w.Write("UpdateTradePacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerId); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Type", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.Type); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Size", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Size); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Trader Tier", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.TraderTier); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Entity Unique Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.EntityUniqueId); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Last Trading Player", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.LastTradingPlayer); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Display Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.DisplayName); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Use New Trade Screen", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.UseNewTradeScreen); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Using Economy Trade", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.UsingEconomyTrade); err != nil {
		return err
	}
	if err := w.Write("UpdateTradePacket.Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.Data); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateTrade(r Decoder) (UpdateTrade, error) {
	var p UpdateTrade
	{
		raw, err := r.Read("UpdateTradePacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Container Id has unexpected decoded type %T", raw)
		}
		p.ContainerId = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Type", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Size", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Size has unexpected decoded type %T", raw)
		}
		p.Size = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Trader Tier", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Trader Tier has unexpected decoded type %T", raw)
		}
		p.TraderTier = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Entity Unique Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Entity Unique Id has unexpected decoded type %T", raw)
		}
		p.EntityUniqueId = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Last Trading Player", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Last Trading Player has unexpected decoded type %T", raw)
		}
		p.LastTradingPlayer = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Display Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Display Name has unexpected decoded type %T", raw)
		}
		p.DisplayName = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Use New Trade Screen", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Use New Trade Screen has unexpected decoded type %T", raw)
		}
		p.UseNewTradeScreen = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Using Economy Trade", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Using Economy Trade has unexpected decoded type %T", raw)
		}
		p.UsingEconomyTrade = value
	}
	{
		raw, err := r.Read("UpdateTradePacket.Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field UpdateTradePacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	return p, nil
}
