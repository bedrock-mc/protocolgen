// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemCompoundTag struct {
	Type  DataItemType
	Value []byte
}

func (*DataItemCompoundTag) isDataItemEntryValue() {}

// Marshal reads or writes DataItemCompoundTag using its canonical wire layout.
func (x *DataItemCompoundTag) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.NBT(&x.Value)
}
