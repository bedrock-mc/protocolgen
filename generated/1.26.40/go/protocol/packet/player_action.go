// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// PlayerAction is sent by the client when it executes any action, for example starting to sprint,
// swim, starting the breaking of a block, dropping an item, etc.
type PlayerAction struct {
	PlayerRuntimeID uint64
	Action          protocol.PlayerActionType
	// BlockPosition is the position of the target block, if the action with the ActionType set
	// concerned a block. If that is not the case, the block position will be zero.
	BlockPosition protocol.BlockPos
	ResultPos     protocol.BlockPos
	Face          int32
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
