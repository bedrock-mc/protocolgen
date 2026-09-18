// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// MobArmorEquipment is sent by the server to the client to update the armour an entity is wearing. It is sent
// for both players and other entities, such as zombies.
type MobArmorEquipment struct {
	TargetRuntimeID uint64
	Head            protocol.NetworkItemStackDescriptorSerializedData
	Torso           protocol.NetworkItemStackDescriptorSerializedData
	Legs            protocol.NetworkItemStackDescriptorSerializedData
	Feet            protocol.NetworkItemStackDescriptorSerializedData
	// Body is the item worn on the body of the entity. Items not wearable on the body will not be rendered.
	Body protocol.NetworkItemStackDescriptorSerializedData
}

// ID returns the protocol ID for MobArmorEquipment.
func (*MobArmorEquipment) ID() uint32 { return IDMobArmorEquipment }

// Marshal reads or writes MobArmorEquipment using its canonical wire layout.
func (pk *MobArmorEquipment) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	pk.Head.Marshal(io)
	pk.Torso.Marshal(io)
	pk.Legs.Marshal(io)
	pk.Feet.Marshal(io)
	pk.Body.Marshal(io)
}
