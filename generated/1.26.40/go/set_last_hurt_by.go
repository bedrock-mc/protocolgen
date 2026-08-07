// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetLastHurtBy struct {
	LastHurtBy ActorType
}

// Marshal reads or writes SetLastHurtBy using its canonical wire layout.
func (x *SetLastHurtBy) Marshal(io IO) {
	IntegerFunc(&x.LastHurtBy, io.Varint32)
}
