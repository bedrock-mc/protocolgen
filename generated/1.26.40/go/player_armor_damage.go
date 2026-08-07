// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []ArmorSlotAndDamagePair
}

// Marshal reads or writes PlayerArmorDamage using its canonical wire layout.
func (x *PlayerArmorDamage) Marshal(io IO) {
	FuncSlice(io, &x.ArmorSlotAndDamagePairs, io.Varuint32, func(value *ArmorSlotAndDamagePair) {
		value.Marshal(io)
	})
}
