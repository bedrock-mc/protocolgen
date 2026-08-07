// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemPos struct {
	Type  DataItemType
	Value BlockPos
}

func (*DataItemPos) isDataItemEntryValue() {}

// Marshal reads or writes DataItemPos using its canonical wire layout.
func (x *DataItemPos) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	x.Value.Marshal(io)
}
