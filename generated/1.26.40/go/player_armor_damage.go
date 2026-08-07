// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []ArmorSlotAndDamagePair
}

// Marshal reads or writes PlayerArmorDamage using its canonical wire layout.
func (x *PlayerArmorDamage) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ArmorSlotAndDamagePairs)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ArmorSlotAndDamagePairs), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ArmorSlotAndDamagePairs))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ArmorSlotAndDamagePairs = make([]ArmorSlotAndDamagePair, int(count1))
	}
	for index2 := range x.ArmorSlotAndDamagePairs {
		x.ArmorSlotAndDamagePairs[index2].Marshal(io)
	}
}
