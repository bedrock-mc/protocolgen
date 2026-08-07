// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundControlSchemeSet struct {
	ControlScheme ControlSchemeScheme
}

// Marshal reads or writes ClientboundControlSchemeSet using its canonical wire layout.
func (x *ClientboundControlSchemeSet) Marshal(io IO) {
	enumValue1 := uint8(x.ControlScheme)
	io.Uint8(&enumValue1)
	x.ControlScheme = ControlSchemeScheme(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
