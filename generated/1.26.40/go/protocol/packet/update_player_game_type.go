// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdatePlayerGameType struct {
	PlayerGameType protocol.GameType
	TargetPlayer   int64
	Tick           uint64
}

// Marshal reads or writes UpdatePlayerGameType using its canonical wire layout.
func (x *UpdatePlayerGameType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PlayerGameType, io.Varint32)
	io.ActorUniqueID(&x.TargetPlayer)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for UpdatePlayerGameType.
func (*UpdatePlayerGameType) ID() uint32 { return IDUpdatePlayerGameType }
