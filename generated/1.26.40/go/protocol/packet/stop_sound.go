// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type StopSound struct {
	SoundName       string
	StopAllSounds   bool
	StopMusicLegacy bool
}

// Marshal reads or writes StopSound using its canonical wire layout.
func (x *StopSound) Marshal(io protocol.IO) {
	io.String(&x.SoundName)
	io.Bool(&x.StopAllSounds)
	io.Bool(&x.StopMusicLegacy)
}

// ID returns the protocol ID for StopSound.
func (*StopSound) ID() uint32 { return IDStopSound }
