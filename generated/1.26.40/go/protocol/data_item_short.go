// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemShort struct {
	Type  DataItemType
	Value int16
}

func (DataItemShort) isDataItemEntryValue() {}

// Marshal reads or writes DataItemShort using its canonical wire layout.
func (x *DataItemShort) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Int16(&x.Value)
}
