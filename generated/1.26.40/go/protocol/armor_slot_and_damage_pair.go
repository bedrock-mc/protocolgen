// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ArmorSlotAndDamagePair struct {
	ArmorSlot LegacyArmorSlot
	Damage    int16
}

// Marshal reads or writes ArmorSlotAndDamagePair using its canonical wire layout.
func (x *ArmorSlotAndDamagePair) Marshal(io IO) {
	IntegerFunc(&x.ArmorSlot, io.Varint32)
	io.Int16(&x.Damage)
}
