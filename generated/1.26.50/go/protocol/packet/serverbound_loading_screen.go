// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType protocol.ServerboundLoadingScreenType
	LoadingScreenID         protocol.Optional[uint32]
}

// ID ...
func (*ServerboundLoadingScreen) ID() uint32 {
	return IDServerboundLoadingScreen
}

func (pk *ServerboundLoadingScreen) Marshal(io protocol.IO) {
	pk.LoadingScreenPacketType.Marshal(io)
	protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
}
