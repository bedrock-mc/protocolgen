// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// Emote is sent by both the server and the client. When the client sends an emote, it sends this
// packet to the server, after which the server will broadcast the packet to other players online.
type Emote struct {
	ActorRuntimeID uint64
	// EmoteID is the ID of the emote to send.
	EmoteID          string
	EmoteLengthTicks uint32
	// Xuid is the Xbox User ID of the player that sent the emote. It is only set when the emote is used
	// by a player that is authenticated with Xbox Live.
	Xuid string
	// PlatformID is an identifier only set for particular platforms when using an emote (presumably
	// only for Nintendo Switch). It is otherwise an empty string, and is used to decide which players
	// are able to emote with each other.
	PlatformID string
	// Flags is a combination of flags that change the way the Emote packet operates. When the server
	// sends this packet to other players, EmoteFlagServerSide must be present.
	Flags uint8
}

// Marshal reads or writes Emote using its canonical wire layout.
func (x *Emote) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.ActorRuntimeID)
	io.String(&x.EmoteID)
	io.Varuint32(&x.EmoteLengthTicks)
	protocol.Minimum(io, &x.EmoteLengthTicks, 0)
	io.String(&x.Xuid)
	io.String(&x.PlatformID)
	io.Uint8(&x.Flags)
	protocol.Minimum(io, &x.Flags, 0)
	protocol.Maximum(io, &x.Flags, 255)
}

// ID returns the protocol ID for Emote.
func (*Emote) ID() uint32 { return IDEmote }
