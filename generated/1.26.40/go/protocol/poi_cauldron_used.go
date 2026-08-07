// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type POICauldronUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemID               int32
}

func (*POICauldronUsed) isEvent() {}

// Marshal reads or writes POICauldronUsed using its canonical wire layout.
func (x *POICauldronUsed) Marshal(io IO) {
	IntegerFunc(&x.BlockInteractionType, io.Uint8)
	io.Varint32(&x.ItemID)
}
