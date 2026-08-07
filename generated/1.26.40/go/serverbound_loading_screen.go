// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType ServerboundLoadingScreenPacketType
	LoadingScreenId         Optional[uint32]
}

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (x *ServerboundLoadingScreen) Marshal(io IO) {
	IntegerFunc(&x.LoadingScreenPacketType, io.Varint32)
	OptionalFunc(io, &x.LoadingScreenId, io.Uint32)
}
