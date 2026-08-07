// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemInt struct {
	Type  DataItemType
	Value int32
}

func (DataItemInt) isDataItemEntryValue() {}

// Marshal reads or writes DataItemInt using its canonical wire layout.
func (x *DataItemInt) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Varint32(&x.Value)
}
