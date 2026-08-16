// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType protocol.ServerboundLoadingScreenType
	LoadingScreenID         protocol.Optional[uint32]
}

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (x *ServerboundLoadingScreen) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.LoadingScreenPacketType, io.Varint32)
	protocol.OptionalFunc(io, &x.LoadingScreenID, io.Uint32)
}

// ID returns the protocol ID for ServerboundLoadingScreen.
func (*ServerboundLoadingScreen) ID() uint32 { return IDServerboundLoadingScreen }
