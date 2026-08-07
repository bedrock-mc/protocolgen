// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerLocation struct {
	TargetActorID int64
	Location      PlayerLocationLocation
}

// Marshal reads or writes PlayerLocation using its canonical wire layout.
func (x *PlayerLocation) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetActorID)
	marshalPlayerLocationLocation(io, &x.Location)
}
