// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MapCreateLockedCopy struct {
	OriginalMapId int64
	NewMapId      int64
}

// Marshal reads or writes MapCreateLockedCopy using its canonical wire layout.
func (x *MapCreateLockedCopy) Marshal(io IO) {
	io.ActorUniqueID(&x.OriginalMapId)
	io.ActorUniqueID(&x.NewMapId)
}
