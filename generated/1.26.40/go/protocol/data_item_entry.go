// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemEntry struct {
	ID      uint32
	Payload DataItemEntryValue
}

// Marshal reads or writes DataItemEntry using its canonical wire layout.
func (x *DataItemEntry) Marshal(io IO) {
	io.Varuint32(&x.ID)
	MarshalDataItemEntryValue(io, &x.Payload)
}
