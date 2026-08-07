// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Emote struct {
	ActorRuntimeId   uint64
	EmoteId          string
	EmoteLengthTicks uint32
	Xuid             string
	PlatformId       string
	Flags            uint8
}

// Marshal reads or writes Emote using its canonical wire layout.
func (x *Emote) Marshal(io IO) {
	io.ActorRuntimeID(&x.ActorRuntimeId)
	io.String(&x.EmoteId)
	io.Varuint32(&x.EmoteLengthTicks)
	io.String(&x.Xuid)
	io.String(&x.PlatformId)
	io.Uint8(&x.Flags)
}
