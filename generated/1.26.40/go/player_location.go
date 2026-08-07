// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerLocation struct {
	TargetActorID ActorUniqueID
	Location      PlayerLocationLocation
}

// Marshal reads or writes PlayerLocation using its canonical wire layout.
func (x *PlayerLocation) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	marshalPlayerLocationLocation(io, &x.Location)
}
