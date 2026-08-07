// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AddVolumeEntity struct {
	EntityNetworkID protocol.EntityNetID
	Components      []byte
	JSONIdentifier  string
	InstanceName    string
	MinBounds       protocol.BlockPos
	MaxBounds       protocol.BlockPos
	DimensionType   protocol.DimensionType
	EngineVersion   string
}

// Marshal reads or writes AddVolumeEntity using its canonical wire layout.
func (x *AddVolumeEntity) Marshal(io protocol.IO) {
	x.EntityNetworkID.Marshal(io)
	io.NBT(&x.Components, protocol.NBTNetwork)
	io.String(&x.JSONIdentifier)
	io.String(&x.InstanceName)
	x.MinBounds.Marshal(io)
	x.MaxBounds.Marshal(io)
	x.DimensionType.Marshal(io)
	io.String(&x.EngineVersion)
}

// ID returns the protocol ID for AddVolumeEntity.
func (*AddVolumeEntity) ID() uint32 { return IDAddVolumeEntity }
