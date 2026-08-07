// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MoveActorAbsolute struct {
	MoveData MoveActorAbsoluteData
}

// Marshal reads or writes MoveActorAbsolute using its canonical wire layout.
func (x *MoveActorAbsolute) Marshal(io IO) {
	x.MoveData.Marshal(io)
}
