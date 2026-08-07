// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerWaxedOrUnwaxedCopper struct {
	PlayerWaxedOrUnwaxedCopperBlockID int32
}

func (*PlayerWaxedOrUnwaxedCopper) isEvent() {}

// Marshal reads or writes PlayerWaxedOrUnwaxedCopper using its canonical wire layout.
func (x *PlayerWaxedOrUnwaxedCopper) Marshal(io IO) {
	io.Varint32(&x.PlayerWaxedOrUnwaxedCopperBlockID)
}
