// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerUpdateEntityOverrides struct {
	TargetID      int64
	PropertyIndex uint32
	Update        PlayerUpdateEntityOverridesUpdate
}

// Marshal reads or writes PlayerUpdateEntityOverrides using its canonical wire layout.
func (x *PlayerUpdateEntityOverrides) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetID)
	io.Varuint32(&x.PropertyIndex)
	marshalPlayerUpdateEntityOverridesUpdate(io, &x.Update)
}
