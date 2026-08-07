// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PrimitiveShapes struct {
	ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved []protocol.PrimitiveShapeData
}

// Marshal reads or writes PrimitiveShapes using its canonical wire layout.
func (x *PrimitiveShapes) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved, io.Varuint32, func(value *protocol.PrimitiveShapeData) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for PrimitiveShapes.
func (*PrimitiveShapes) ID() uint32 { return IDPrimitiveShapes }
