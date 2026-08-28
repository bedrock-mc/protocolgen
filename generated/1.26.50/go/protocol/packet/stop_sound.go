// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// StopSound is sent by the server to stop a sound playing to the player, such as a playing music
// disk track or other long-lasting sounds.
type StopSound struct {
	// SoundName is the name of the sound that should be stopped from playing. If no sound with this
	// name is currently active, the packet is ignored.
	SoundName string
	// StopAllSounds specifies if all sounds currently playing to the player should be stopped. If set
	// to true, the SoundName field may be left empty.
	StopAllSounds bool
	// StopMusicLegacy is currently unknown.
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
