// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateClientOptions struct {
	GraphicsModeChange    Optional[GraphicsMode]
	FilterProfanityChange Optional[bool]
}

// Marshal reads or writes UpdateClientOptions using its canonical wire layout.
func (x *UpdateClientOptions) Marshal(io IO) {
	OptionalFunc(io, &x.GraphicsModeChange, func(value *GraphicsMode) {
		IntegerFunc(value, io.Uint8)
	})
	OptionalFunc(io, &x.FilterProfanityChange, io.Bool)
}
