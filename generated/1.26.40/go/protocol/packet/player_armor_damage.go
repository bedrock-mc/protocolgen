// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []protocol.ArmorSlotAndDamagePair
}

// Marshal reads or writes PlayerArmorDamage using its canonical wire layout.
func (x *PlayerArmorDamage) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.ArmorSlotAndDamagePairs, io.Varuint32, func(value *protocol.ArmorSlotAndDamagePair) {
		value.Marshal(io)
	})
}
