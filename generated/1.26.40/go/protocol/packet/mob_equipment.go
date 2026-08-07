// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MobEquipment struct {
	TargetRuntimeID uint64
	Item            protocol.CerealizerNetworkItemStackDescriptorSerializedData
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
