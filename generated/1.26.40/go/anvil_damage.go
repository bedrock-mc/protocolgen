// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AnvilDamage struct {
	BlockPosition BlockPos
}

// Marshal reads or writes AnvilDamage using its canonical wire layout.
func (x *AnvilDamage) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
}
