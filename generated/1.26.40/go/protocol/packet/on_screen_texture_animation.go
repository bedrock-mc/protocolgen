// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type OnScreenTextureAnimation struct {
	EffectID uint32
}

// Marshal reads or writes OnScreenTextureAnimation using its canonical wire layout.
func (x *OnScreenTextureAnimation) Marshal(io protocol.IO) {
	io.Uint32(&x.EffectID)
}

// ID returns the protocol ID for OnScreenTextureAnimation.
func (*OnScreenTextureAnimation) ID() uint32 { return IDOnScreenTextureAnimation }
