// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundControlSchemeSet struct {
	ControlScheme ControlSchemeScheme
}

// Marshal reads or writes ClientboundControlSchemeSet using its canonical wire layout.
func (x *ClientboundControlSchemeSet) Marshal(io IO) {
	IntegerFunc(&x.ControlScheme, io.Uint8)
}
