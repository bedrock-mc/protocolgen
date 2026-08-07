// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type NetworkSettings struct {
	CompressionThreshold    uint16
	CompressionAlgorithm    protocol.PacketCompressionAlgorithm
	ClientThrottleEnabled   bool
	ClientThrottleThreshold uint8
	ClientThrottleScalar    float32
}

// Marshal reads or writes NetworkSettings using its canonical wire layout.
func (x *NetworkSettings) Marshal(io protocol.IO) {
	io.Uint16(&x.CompressionThreshold)
	protocol.IntegerFunc(&x.CompressionAlgorithm, io.Uint16)
	io.Bool(&x.ClientThrottleEnabled)
	io.Uint8(&x.ClientThrottleThreshold)
	io.Float32(&x.ClientThrottleScalar)
}

// ID returns the protocol ID for NetworkSettings.
func (*NetworkSettings) ID() uint32 { return IDNetworkSettings }
