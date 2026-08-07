// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MoveActorDelta struct {
	MoveData MoveActorDeltaData
}

// Marshal reads or writes MoveActorDelta using its canonical wire layout.
func (x *MoveActorDelta) Marshal(io IO) {
	x.MoveData.Marshal(io)
}
