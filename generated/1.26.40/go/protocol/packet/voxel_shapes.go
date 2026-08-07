// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type VoxelShapes struct {
	Shapes           []protocol.VoxelShapesSerializableVoxelShape
	NameMap          []protocol.OrderedEntry[string, protocol.VoxelShapesRegistryHandle]
	CustomShapeCount uint16
}

// Marshal reads or writes VoxelShapes using its canonical wire layout.
func (x *VoxelShapes) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Shapes)
	protocol.OrderedMap(io, &x.NameMap, io.Varuint32, io.String, func(value *protocol.VoxelShapesRegistryHandle) {
		value.Marshal(io)
	})
	io.Uint16(&x.CustomShapeCount)
}

// ID returns the protocol ID for VoxelShapes.
func (*VoxelShapes) ID() uint32 { return IDVoxelShapes }
