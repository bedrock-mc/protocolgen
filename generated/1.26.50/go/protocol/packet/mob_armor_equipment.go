// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

type MobArmorEquipment struct {
	TargetRuntimeID uint64
	Head            protocol.NetworkItemStackDescriptorSerializedData
	Torso           protocol.NetworkItemStackDescriptorSerializedData
	Legs            protocol.NetworkItemStackDescriptorSerializedData
	Feet            protocol.NetworkItemStackDescriptorSerializedData
	Body            protocol.NetworkItemStackDescriptorSerializedData
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

// ID returns the protocol ID for MobArmorEquipment.
func (*MobArmorEquipment) ID() uint32 { return IDMobArmorEquipment }
