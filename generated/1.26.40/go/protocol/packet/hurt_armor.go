// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type HurtArmor struct {
	Cause      int32
	Damage     int32
	ArmorSlots uint64
}

// ID returns the protocol ID for HurtArmor.
func (*HurtArmor) ID() uint32 { return IDHurtArmor }

// Marshal reads or writes HurtArmor using its canonical wire layout.
func (pk *HurtArmor) Marshal(io protocol.IO) {
	io.Varint32(&pk.Cause)
	io.Varint32(&pk.Damage)
	io.Varuint64(&pk.ArmorSlots)
}
