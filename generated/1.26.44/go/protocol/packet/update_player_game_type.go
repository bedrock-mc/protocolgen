// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// UpdatePlayerGameType is sent by the server to change the game mode of a player. It is
// functionally identical to the SetPlayerGameType packet.
type UpdatePlayerGameType struct {
	PlayerGameType protocol.GameType
	TargetPlayer   int64
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// Marshal reads or writes UpdatePlayerGameType using its canonical wire layout.
func (x *UpdatePlayerGameType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PlayerGameType, io.Varint32)
	io.ActorUniqueID(&x.TargetPlayer)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for UpdatePlayerGameType.
func (*UpdatePlayerGameType) ID() uint32 { return IDUpdatePlayerGameType }
