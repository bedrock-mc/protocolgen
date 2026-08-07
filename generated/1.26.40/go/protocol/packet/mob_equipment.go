// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// MobEquipment is sent by the client to the server and the server to the client to make the other
// side aware of the new item that an entity is holding. It is used to show the item in the hand of
// entities such as zombies too.
type MobEquipment struct {
	TargetRuntimeID uint64
	Item            protocol.NetworkItemStackDescriptorSerializedData
	Slot            uint8
	SelectedSlot    uint8
	ContainerID     uint8
}

// Marshal reads or writes MobEquipment using its canonical wire layout.
func (x *MobEquipment) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.Item.Marshal(io)
	io.Uint8(&x.Slot)
	io.Uint8(&x.SelectedSlot)
	io.Uint8(&x.ContainerID)
}

// ID returns the protocol ID for MobEquipment.
func (*MobEquipment) ID() uint32 { return IDMobEquipment }
