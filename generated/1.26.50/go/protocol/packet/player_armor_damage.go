// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []protocol.ArmorSlotAndDamagePair
}

// ID returns the protocol ID for PlayerArmorDamage.
func (*PlayerArmorDamage) ID() uint32 { return IDPlayerArmorDamage }

// Marshal reads or writes PlayerArmorDamage using its canonical wire layout.
func (pk *PlayerArmorDamage) Marshal(io protocol.IO) {
	protocol.SliceLimits(io, &pk.ArmorSlotAndDamagePairs, 0, 5)
}
