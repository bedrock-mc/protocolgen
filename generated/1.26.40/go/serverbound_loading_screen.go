// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType ServerboundLoadingScreenPacketType
	LoadingScreenId         Optional[uint32]
}

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (x *ServerboundLoadingScreen) Marshal(io IO) {
	enumValue1 := int32(x.LoadingScreenPacketType)
	io.Varint32(&enumValue1)
	x.LoadingScreenPacketType = ServerboundLoadingScreenPacketType(enumValue1)
	switch int64(enumValue1) {
	case 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.LoadingScreenId.set)
	if x.LoadingScreenId.set {
		io.Uint32(&x.LoadingScreenId.val)
	} else if io.Reading() {
		var zero uint32
		x.LoadingScreenId.val = zero
	}
}
