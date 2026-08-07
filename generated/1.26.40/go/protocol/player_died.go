// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerDied struct {
	InstigatorActorID    int32
	InstigatorMobVariant int32
	DamageSource         int32
	DiedInRaid           bool
}

func (*PlayerDied) isEvent() {}

// Marshal reads or writes PlayerDied using its canonical wire layout.
func (x *PlayerDied) Marshal(io IO) {
	io.Varint32(&x.InstigatorActorID)
	io.Varint32(&x.InstigatorMobVariant)
	io.Varint32(&x.DamageSource)
	io.Bool(&x.DiedInRaid)
}
