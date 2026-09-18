// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType protocol.ServerboundLoadingScreenType
	LoadingScreenID         protocol.Optional[uint32]
}

// ID returns the protocol ID for ServerboundLoadingScreen.
func (*ServerboundLoadingScreen) ID() uint32 { return IDServerboundLoadingScreen }

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (pk *ServerboundLoadingScreen) Marshal(io protocol.IO) {
	pk.LoadingScreenPacketType.Marshal(io)
	protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
}
