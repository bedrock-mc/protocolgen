// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerBlockActionData struct {
	PlayerActionType PlayerActionType
	Position         BlockPos
	Facing           int32
}

// Marshal reads or writes PlayerBlockActionData using its canonical wire layout.
func (x *PlayerBlockActionData) Marshal(io IO) {
	IntegerFunc(&x.PlayerActionType, io.Varint32)
	x.Position.Marshal(io)
	io.Varint32(&x.Facing)
}
