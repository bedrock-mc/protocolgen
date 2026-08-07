// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type DataItemVec3 struct {
	Type  DataItemType
	Value mgl32.Vec3
}

func (DataItemVec3) isDataItemEntryValue() {}

// Marshal reads or writes DataItemVec3 using its canonical wire layout.
func (x *DataItemVec3) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Vec3(&x.Value)
}
