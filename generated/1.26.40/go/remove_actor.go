// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RemoveActor struct {
	TargetActorID ActorUniqueID
}

// Marshal reads or writes RemoveActor using its canonical wire layout.
func (x *RemoveActor) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
}
