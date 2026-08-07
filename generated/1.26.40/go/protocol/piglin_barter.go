// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PiglinBarter struct {
	ItemID                      int32
	WasTargetingBarteringPlayer bool
}

func (*PiglinBarter) isEvent() {}

// Marshal reads or writes PiglinBarter using its canonical wire layout.
func (x *PiglinBarter) Marshal(io IO) {
	io.Varint32(&x.ItemID)
	io.Bool(&x.WasTargetingBarteringPlayer)
}
