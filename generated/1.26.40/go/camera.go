// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Camera struct {
	CameraID       int64
	TargetPlayerID int64
}

// Marshal reads or writes Camera using its canonical wire layout.
func (x *Camera) Marshal(io IO) {
	io.ActorUniqueID(&x.CameraID)
	io.ActorUniqueID(&x.TargetPlayerID)
}
