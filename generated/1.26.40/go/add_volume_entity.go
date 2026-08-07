// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AddVolumeEntity struct {
	EntityNetworkId EntityNetId
	Components      []byte
	JSONIdentifier  string
	InstanceName    string
	MinBounds       BlockPos
	MaxBounds       BlockPos
	DimensionType   DimensionType
	EngineVersion   string
}

// Marshal reads or writes AddVolumeEntity using its canonical wire layout.
func (x *AddVolumeEntity) Marshal(io IO) {
	x.EntityNetworkId.Marshal(io)
	io.NBT(&x.Components)
	io.String(&x.JSONIdentifier)
	io.String(&x.InstanceName)
	x.MinBounds.Marshal(io)
	x.MaxBounds.Marshal(io)
	x.DimensionType.Marshal(io)
	io.String(&x.EngineVersion)
}
