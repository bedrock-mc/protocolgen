// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateClientOptions struct {
	GraphicsModeChange    Optional[GraphicsMode]
	FilterProfanityChange Optional[bool]
}

// Marshal reads or writes UpdateClientOptions using its canonical wire layout.
func (x *UpdateClientOptions) Marshal(io IO) {
	io.Bool(&x.GraphicsModeChange.set)
	if x.GraphicsModeChange.set {
		enumValue1 := uint8(x.GraphicsModeChange.val)
		io.Uint8(&enumValue1)
		x.GraphicsModeChange.val = GraphicsMode(enumValue1)
		switch int64(enumValue1) {
		case 0, 1, 2, 3:
		default:
			io.InvalidValue(enumValue1, "unknown enum value")
		}
	} else if io.Reading() {
		var zero GraphicsMode
		x.GraphicsModeChange.val = zero
	}
	io.Bool(&x.FilterProfanityChange.set)
	if x.FilterProfanityChange.set {
		io.Bool(&x.FilterProfanityChange.val)
	} else if io.Reading() {
		var zero bool
		x.FilterProfanityChange.val = zero
	}
}
