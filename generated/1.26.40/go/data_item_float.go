// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DataItemFloat struct {
	Type  DataItemType
	Value float32
}

func (DataItemFloat) isDataItemEntryValue() {}

// Marshal reads or writes DataItemFloat using its canonical wire layout.
func (x *DataItemFloat) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.Float32(&x.Value)
}
