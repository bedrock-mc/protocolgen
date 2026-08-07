// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RemoveActor struct {
	TargetActorID int64
}

// Marshal reads or writes RemoveActor using its canonical wire layout.
func (x *RemoveActor) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetActorID)
}
