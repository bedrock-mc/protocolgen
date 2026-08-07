// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateBlockSynced struct {
	BlockPosition    BlockPos
	BlockRuntimeID   uint32
	Flags            uint32
	Layer            uint32
	UniqueActorId    uint64
	ActorSyncMessage uint64
}

func (p *UpdateBlockSynced) Encode(w Encoder) error {
	if err := w.Write("UpdateBlockSyncedPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockSyncedPacket.Block Runtime ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.BlockRuntimeID); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockSyncedPacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Flags); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockSyncedPacket.Layer", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Layer); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockSyncedPacket.Unique Actor Id", Shape{Kind: "primitive", PrimitiveCode: "var_u64"}, p.UniqueActorId); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockSyncedPacket.Actor Sync Message", Shape{Kind: "primitive", PrimitiveCode: "var_u64"}, p.ActorSyncMessage); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateBlockSynced(r Decoder) (UpdateBlockSynced, error) {
	var p UpdateBlockSynced
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Block Runtime ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Block Runtime ID has unexpected decoded type %T", raw)
		}
		p.BlockRuntimeID = value
	}
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Flags has unexpected decoded type %T", raw)
		}
		p.Flags = value
	}
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Layer", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Layer has unexpected decoded type %T", raw)
		}
		p.Layer = value
	}
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Unique Actor Id", Shape{Kind: "primitive", PrimitiveCode: "var_u64"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Unique Actor Id has unexpected decoded type %T", raw)
		}
		p.UniqueActorId = value
	}
	{
		raw, err := r.Read("UpdateBlockSyncedPacket.Actor Sync Message", Shape{Kind: "primitive", PrimitiveCode: "var_u64"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockSyncedPacket.Actor Sync Message has unexpected decoded type %T", raw)
		}
		p.ActorSyncMessage = value
	}
	return p, nil
}
