// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type OnScreenTextureAnimation struct {
	EffectId uint32
}

// Marshal reads or writes OnScreenTextureAnimation using its canonical wire layout.
func (x *OnScreenTextureAnimation) Marshal(io IO) {
	io.Uint32(&x.EffectId)
}
