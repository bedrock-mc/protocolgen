// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerAction struct {
	PlayerRuntimeID ActorRuntimeID
	Action          PlayerActionType
	BlockPosition   BlockPos
	ResultPos       BlockPos
	Face            int32
}

// Marshal reads or writes PlayerAction using its canonical wire layout.
func (x *PlayerAction) Marshal(io IO) {
	x.PlayerRuntimeID.Marshal(io)
	enumValue1 := int32(x.Action)
	io.Varint32(&enumValue1)
	x.Action = PlayerActionType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.BlockPosition.Marshal(io)
	x.ResultPos.Marshal(io)
	io.Varint32(&x.Face)
}
