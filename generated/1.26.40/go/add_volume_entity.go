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
