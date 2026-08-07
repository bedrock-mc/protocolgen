// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MapCreateLockedCopy struct {
	OriginalMapId ActorUniqueID
	NewMapId      ActorUniqueID
}

// Marshal reads or writes MapCreateLockedCopy using its canonical wire layout.
func (x *MapCreateLockedCopy) Marshal(io IO) {
	x.OriginalMapId.Marshal(io)
	x.NewMapId.Marshal(io)
}
