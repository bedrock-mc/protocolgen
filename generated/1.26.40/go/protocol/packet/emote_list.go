// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

type EmoteList struct {
	RuntimeID     uint64
	EmotePieceIds []uuid.UUID
}

// Marshal reads or writes EmoteList using its canonical wire layout.
func (x *EmoteList) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.RuntimeID)
	protocol.FuncSlice(io, &x.EmotePieceIds, io.Varuint32, io.UUID)
}

// ID returns the protocol ID for EmoteList.
func (*EmoteList) ID() uint32 { return IDEmoteList }
