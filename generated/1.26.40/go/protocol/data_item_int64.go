// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemInt64 struct {
	Type  DataItemType
	Value int64
}

func (*DataItemInt64) isDataItemEntryValue() {}

// Marshal reads or writes DataItemInt64 using its canonical wire layout.
func (x *DataItemInt64) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Varint64(&x.Value)
}
