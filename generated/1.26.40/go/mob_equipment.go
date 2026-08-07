// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MobEquipment struct {
	TargetRuntimeID ActorRuntimeID
	Item            CerealizerNetworkItemStackDescriptorSerializedData
	Slot            uint8
	SelectedSlot    uint8
	ContainerID     uint8
}

// Marshal reads or writes MobEquipment using its canonical wire layout.
func (x *MobEquipment) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	x.Item.Marshal(io)
	io.Uint8(&x.Slot)
	io.Uint8(&x.SelectedSlot)
	io.Uint8(&x.ContainerID)
}
