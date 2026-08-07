// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CompletedUsingItem struct {
	ItemId        int16
	ItemUseMethod int32
}

// Marshal reads or writes CompletedUsingItem using its canonical wire layout.
func (x *CompletedUsingItem) Marshal(io IO) {
	io.Int16(&x.ItemId)
	io.Int32(&x.ItemUseMethod)
}
