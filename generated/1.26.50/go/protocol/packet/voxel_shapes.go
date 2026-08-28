// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// VoxelShapes is sent by the server to send voxel shape data to the client.
type VoxelShapes struct {
	// Shapes is a list of voxel shapes.
	Shapes []protocol.VoxelShapesSerializableVoxelShape
	// NameMap is a map of shape names to IDs.
	NameMap []protocol.OrderedEntry[string, protocol.VoxelShapesRegistryHandle]
	// CustomShapeCount is the number of custom shapes.
	CustomShapeCount uint16
}

// Marshal reads or writes VoxelShapes using its canonical wire layout.
func (x *VoxelShapes) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Shapes)
	protocol.OrderedMap(io, &x.NameMap, io.Varuint32, io.String, func(value *protocol.VoxelShapesRegistryHandle) {
		value.Marshal(io)
	})
	io.Uint16(&x.CustomShapeCount)
	protocol.Minimum(io, &x.CustomShapeCount, 0)
}

// ID returns the protocol ID for VoxelShapes.
func (*VoxelShapes) ID() uint32 { return IDVoxelShapes }
