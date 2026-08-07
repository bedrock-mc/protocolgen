// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MobArmorEquipment struct {
	TargetRuntimeID uint64
	Head            protocol.CerealizerNetworkItemStackDescriptorSerializedData
	Torso           protocol.CerealizerNetworkItemStackDescriptorSerializedData
	Legs            protocol.CerealizerNetworkItemStackDescriptorSerializedData
	Feet            protocol.CerealizerNetworkItemStackDescriptorSerializedData
	Body            protocol.CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes MobArmorEquipment using its canonical wire layout.
func (x *MobArmorEquipment) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.Head.Marshal(io)
	x.Torso.Marshal(io)
	x.Legs.Marshal(io)
	x.Feet.Marshal(io)
	x.Body.Marshal(io)
}
