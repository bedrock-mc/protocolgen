// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Camera struct {
	CameraID       ActorUniqueID
	TargetPlayerID ActorUniqueID
}

// Marshal reads or writes Camera using its canonical wire layout.
func (x *Camera) Marshal(io IO) {
	x.CameraID.Marshal(io)
	x.TargetPlayerID.Marshal(io)
}
