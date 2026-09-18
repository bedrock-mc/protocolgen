// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type MobArmorEquipment struct {
	TargetRuntimeID uint64
	Head            protocol.NetworkItemStackDescriptorSerializedData
	Torso           protocol.NetworkItemStackDescriptorSerializedData
	Legs            protocol.NetworkItemStackDescriptorSerializedData
	Feet            protocol.NetworkItemStackDescriptorSerializedData
	Body            protocol.NetworkItemStackDescriptorSerializedData
}

// ID ...
func (*MobArmorEquipment) ID() uint32 {
	return IDMobArmorEquipment
}

func (pk *MobArmorEquipment) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	pk.Head.Marshal(io)
	pk.Torso.Marshal(io)
	pk.Legs.Marshal(io)
	pk.Feet.Marshal(io)
	pk.Body.Marshal(io)
}
