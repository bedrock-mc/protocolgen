// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type DimensionDefinition struct {
	HeightMaximum int32
	HeightMinimum int32
	GeneratorType GeneratorType
	DimensionType DimensionType
	PackId        uuid.UUID
}

// Marshal reads or writes DimensionDefinition using its canonical wire layout.
func (x *DimensionDefinition) Marshal(io IO) {
	io.Varint32(&x.HeightMaximum)
	io.Varint32(&x.HeightMinimum)
	IntegerFunc(&x.GeneratorType, io.Varint32)
	x.DimensionType.Marshal(io)
	io.UUID(&x.PackId)
}
