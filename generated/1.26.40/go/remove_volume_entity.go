// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RemoveVolumeEntity struct {
	EntityNetworkId EntityNetId
	DimensionType   DimensionType
}

// Marshal reads or writes RemoveVolumeEntity using its canonical wire layout.
func (x *RemoveVolumeEntity) Marshal(io IO) {
	x.EntityNetworkId.Marshal(io)
	x.DimensionType.Marshal(io)
}
