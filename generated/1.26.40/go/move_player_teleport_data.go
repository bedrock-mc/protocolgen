// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MovePlayerTeleportData struct {
	TeleportationCause int32
	SourceActorType    int32
}

// Marshal reads or writes MovePlayerTeleportData using its canonical wire layout.
func (x *MovePlayerTeleportData) Marshal(io IO) {
	io.Int32(&x.TeleportationCause)
	io.Int32(&x.SourceActorType)
}
