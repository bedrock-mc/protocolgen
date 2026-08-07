// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemData struct {
	ItemName          string
	ItemID            int16
	IsComponentBased  bool
	ItemVersion       ItemVersion
	ItemComponentData []byte
}

// Marshal reads or writes ItemData using its canonical wire layout.
func (x *ItemData) Marshal(io IO) {
	io.String(&x.ItemName)
	io.Int16(&x.ItemID)
	io.Bool(&x.IsComponentBased)
	IntegerFunc(&x.ItemVersion, io.Varint32)
	io.NBT(&x.ItemComponentData)
}
