// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ComposterUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemID               int32
}

func (*ComposterUsed) isEvent() {}

// Marshal reads or writes ComposterUsed using its canonical wire layout.
func (x *ComposterUsed) Marshal(io IO) {
	IntegerFunc(&x.BlockInteractionType, io.Uint8)
	io.Varint32(&x.ItemID)
}
