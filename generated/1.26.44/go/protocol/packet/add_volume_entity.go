// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// AddVolumeEntity sends a volume entity's definition and metadata from server to client.
type AddVolumeEntity struct {
	EntityNetworkID protocol.EntityNetID
	Components      []byte
	JSONIdentifier  string
	InstanceName    string
	MinBounds       protocol.BlockPos
	MaxBounds       protocol.BlockPos
	DimensionType   protocol.DimensionType
	// EngineVersion is the engine version the entity is using, for example, '1.17.0'.
	EngineVersion string
}

// Marshal reads or writes AddVolumeEntity using its canonical wire layout.
func (x *AddVolumeEntity) Marshal(io protocol.IO) {
	x.EntityNetworkID.Marshal(io)
	io.NBT(&x.Components, protocol.NBTNetwork)
	io.StringLimits(&x.JSONIdentifier, 1, 18446744073709551615)
	io.StringLimits(&x.InstanceName, 1, 18446744073709551615)
	x.MinBounds.Marshal(io)
	x.MaxBounds.Marshal(io)
	x.DimensionType.Marshal(io)
	io.String(&x.EngineVersion)
}

// ID returns the protocol ID for AddVolumeEntity.
func (*AddVolumeEntity) ID() uint32 { return IDAddVolumeEntity }
