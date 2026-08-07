// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerAction struct {
	PlayerRuntimeID uint64
	Action          protocol.PlayerActionType
	BlockPosition   protocol.BlockPos
	ResultPos       protocol.BlockPos
	Face            int32
}

// Marshal reads or writes PlayerAction using its canonical wire layout.
func (x *PlayerAction) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.PlayerRuntimeID)
	protocol.IntegerFunc(&x.Action, io.Varint32)
	x.BlockPosition.Marshal(io)
	x.ResultPos.Marshal(io)
	io.Varint32(&x.Face)
}

// ID returns the protocol ID for PlayerAction.
func (*PlayerAction) ID() uint32 { return IDPlayerAction }
