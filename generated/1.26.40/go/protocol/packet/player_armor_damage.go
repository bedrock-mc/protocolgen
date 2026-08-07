// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []protocol.ArmorSlotAndDamagePair
}

// Marshal reads or writes PlayerArmorDamage using its canonical wire layout.
func (x *PlayerArmorDamage) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.ArmorSlotAndDamagePairs)
}

// ID returns the protocol ID for PlayerArmorDamage.
func (*PlayerArmorDamage) ID() uint32 { return IDPlayerArmorDamage }
