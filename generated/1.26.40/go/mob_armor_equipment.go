// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MobArmorEquipment struct {
	TargetRuntimeID uint64
	Head            CerealizerNetworkItemStackDescriptorSerializedData
	Torso           CerealizerNetworkItemStackDescriptorSerializedData
	Legs            CerealizerNetworkItemStackDescriptorSerializedData
	Feet            CerealizerNetworkItemStackDescriptorSerializedData
	Body            CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes MobArmorEquipment using its canonical wire layout.
func (x *MobArmorEquipment) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.Head.Marshal(io)
	x.Torso.Marshal(io)
	x.Legs.Marshal(io)
	x.Feet.Marshal(io)
	x.Body.Marshal(io)
}
