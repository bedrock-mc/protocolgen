// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/google/uuid"
)

type EmoteList struct {
	RuntimeId     ActorRuntimeID
	EmotePieceIds []uuid.UUID
}

func (p *EmoteList) Encode(w Encoder) error {
	if err := w.Write("EmoteListPacket.Runtime id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.RuntimeId); err != nil {
		return err
	}
	if err := w.Write("EmoteListPacket.Emote piece ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, p.EmotePieceIds); err != nil {
		return err
	}
	return nil
}

func DecodeEmoteList(r Decoder) (EmoteList, error) {
	var p EmoteList
	{
		raw, err := r.Read("EmoteListPacket.Runtime id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field EmoteListPacket.Runtime id has unexpected decoded type %T", raw)
		}
		p.RuntimeId = value
	}
	{
		raw, err := r.Read("EmoteListPacket.Emote piece ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]uuid.UUID)
		if !ok {
			return p, fmt.Errorf("field EmoteListPacket.Emote piece ids has unexpected decoded type %T", raw)
		}
		p.EmotePieceIds = value
	}
	return p, nil
}
