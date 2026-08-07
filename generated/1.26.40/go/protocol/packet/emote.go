// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Emote struct {
	ActorRuntimeID   uint64
	EmoteID          string
	EmoteLengthTicks uint32
	Xuid             string
	PlatformID       string
	Flags            uint8
}

// Marshal reads or writes Emote using its canonical wire layout.
func (x *Emote) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.ActorRuntimeID)
	io.String(&x.EmoteID)
	io.Varuint32(&x.EmoteLengthTicks)
	io.String(&x.Xuid)
	io.String(&x.PlatformID)
	io.Uint8(&x.Flags)
}

// ID returns the protocol ID for Emote.
func (*Emote) ID() uint32 { return IDEmote }
