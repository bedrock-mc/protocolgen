// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MaterialReducerEntryOutput struct {
	ItemID    int32
	ItemCount int32
}

// Marshal reads or writes MaterialReducerEntryOutput using its canonical wire layout.
func (x *MaterialReducerEntryOutput) Marshal(io IO) {
	io.Varint32(&x.ItemID)
	io.Varint32(&x.ItemCount)
}
