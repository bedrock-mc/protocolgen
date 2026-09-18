// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// MobEquipment is sent by the client to the server and the server to the client to make the other side aware
// of the new item that an entity is holding. It is used to show the item in the hand of entities such as
// zombies too.
type MobEquipment struct {
	TargetRuntimeID uint64
	Item            protocol.NetworkItemStackDescriptorSerializedData
	Slot            uint8
	SelectedSlot    uint8
	ContainerID     uint8
}

// ID returns the protocol ID for MobEquipment.
func (*MobEquipment) ID() uint32 { return IDMobEquipment }

// Marshal reads or writes MobEquipment using its canonical wire layout.
func (pk *MobEquipment) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	pk.Item.Marshal(io)
	io.Uint8(&pk.Slot)
	io.Uint8(&pk.SelectedSlot)
	io.Uint8(&pk.ContainerID)
}
