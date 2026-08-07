// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NetworkSettings struct {
	CompressionThreshold    uint16
	CompressionAlgorithm    PacketCompressionAlgorithm
	ClientThrottleEnabled   bool
	ClientThrottleThreshold uint8
	ClientThrottleScalar    float32
}

// Marshal reads or writes NetworkSettings using its canonical wire layout.
func (x *NetworkSettings) Marshal(io IO) {
	io.Uint16(&x.CompressionThreshold)
	enumValue1 := uint16(x.CompressionAlgorithm)
	io.Uint16(&enumValue1)
	x.CompressionAlgorithm = PacketCompressionAlgorithm(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 65535:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.ClientThrottleEnabled)
	io.Uint8(&x.ClientThrottleThreshold)
	io.Float32(&x.ClientThrottleScalar)
}
