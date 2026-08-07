// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type HurtArmor struct {
	Cause      int32
	Damage     int32
	ArmorSlots uint64
}

// Marshal reads or writes HurtArmor using its canonical wire layout.
func (x *HurtArmor) Marshal(io IO) {
	io.Varint32(&x.Cause)
	io.Varint32(&x.Damage)
	io.Varuint64(&x.ArmorSlots)
}
