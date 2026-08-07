// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerAction struct {
	PlayerRuntimeID uint64
	Action          PlayerActionType
	BlockPosition   BlockPos
	ResultPos       BlockPos
	Face            int32
}

// Marshal reads or writes PlayerAction using its canonical wire layout.
func (x *PlayerAction) Marshal(io IO) {
	io.ActorRuntimeID(&x.PlayerRuntimeID)
	IntegerFunc(&x.Action, io.Varint32)
	x.BlockPosition.Marshal(io)
	x.ResultPos.Marshal(io)
	io.Varint32(&x.Face)
}
