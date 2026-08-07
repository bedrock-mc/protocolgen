// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BellUsed struct {
	ItemID int32
}

func (*BellUsed) isEvent() {}

// Marshal reads or writes BellUsed using its canonical wire layout.
func (x *BellUsed) Marshal(io IO) {
	io.Varint32(&x.ItemID)
}
