// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DataItemByte struct {
	Type  DataItemType
	Value int8
}

func (DataItemByte) isDataItemEntryValue() {}

// Marshal reads or writes DataItemByte using its canonical wire layout.
func (x *DataItemByte) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Int8(&x.Value)
}
