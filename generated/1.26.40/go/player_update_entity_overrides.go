// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerUpdateEntityOverrides struct {
	TargetID      ActorUniqueID
	PropertyIndex uint32
	Update        PlayerUpdateEntityOverridesUpdate
}

// Marshal reads or writes PlayerUpdateEntityOverrides using its canonical wire layout.
func (x *PlayerUpdateEntityOverrides) Marshal(io IO) {
	x.TargetID.Marshal(io)
	io.Varuint32(&x.PropertyIndex)
	marshalPlayerUpdateEntityOverridesUpdate(io, &x.Update)
}
