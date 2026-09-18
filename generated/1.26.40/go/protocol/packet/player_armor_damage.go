// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []protocol.ArmorSlotAndDamagePair
}

// ID ...
func (*PlayerArmorDamage) ID() uint32 {
	return IDPlayerArmorDamage
}

func (pk *PlayerArmorDamage) Marshal(io protocol.IO) {
	protocol.SliceLimits(io, &pk.ArmorSlotAndDamagePairs, 0, 5)
}
