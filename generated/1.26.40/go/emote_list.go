// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type EmoteList struct {
	RuntimeId     uint64
	EmotePieceIds []uuid.UUID
}

// Marshal reads or writes EmoteList using its canonical wire layout.
func (x *EmoteList) Marshal(io IO) {
	io.ActorRuntimeID(&x.RuntimeId)
	FuncSlice(io, &x.EmotePieceIds, io.Varuint32, io.UUID)
}
