// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type StopSound struct {
	SoundName       string
	StopAllSounds   bool
	StopMusicLegacy bool
}

// Marshal reads or writes StopSound using its canonical wire layout.
func (x *StopSound) Marshal(io IO) {
	io.String(&x.SoundName)
	io.Bool(&x.StopAllSounds)
	io.Bool(&x.StopMusicLegacy)
}
