// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType protocol.ServerboundLoadingScreenPacketType
	LoadingScreenId         protocol.Optional[uint32]
}

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (x *ServerboundLoadingScreen) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.LoadingScreenPacketType, io.Varint32)
	protocol.OptionalFunc(io, &x.LoadingScreenId, io.Uint32)
}
