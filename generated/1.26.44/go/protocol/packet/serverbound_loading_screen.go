// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ServerboundLoadingScreen is sent by the client to tell the server about the state of the loading screen
// that the client is currently displaying.
type ServerboundLoadingScreen struct {
	// Type is the type of the loading screen event. It is one of the constants that may be found above.
	LoadingScreenPacketType protocol.ServerboundLoadingScreenType
	// LoadingScreenID is the ID of the screen that was previously sent by the server in the ChangeDimension
	// packet. The server should validate that the ID matches the last one it sent.
	LoadingScreenID protocol.Optional[uint32]
}

// ID returns the protocol ID for ServerboundLoadingScreen.
func (*ServerboundLoadingScreen) ID() uint32 { return IDServerboundLoadingScreen }

// Marshal reads or writes ServerboundLoadingScreen using its canonical wire layout.
func (pk *ServerboundLoadingScreen) Marshal(io protocol.IO) {
	pk.LoadingScreenPacketType.Marshal(io)
	protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
}
