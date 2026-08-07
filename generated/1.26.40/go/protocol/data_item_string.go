// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemString struct {
	Type  DataItemType
	Value string
}

func (*DataItemString) isDataItemEntryValue() {}

// Marshal reads or writes DataItemString using its canonical wire layout.
func (x *DataItemString) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	io.String(&x.Value)
}
