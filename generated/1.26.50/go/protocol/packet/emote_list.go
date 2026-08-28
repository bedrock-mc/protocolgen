// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/google/uuid"
)

// EmoteList is sent by the client every time it joins the server and when it equips new emotes. It
// may be used by the server to find out which emotes the client has available. If the player has no
// emotes equipped, this packet is not sent. Under certain circumstances, this packet is also sent
// from the server to the client, but I was unable to find when this is done.
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
