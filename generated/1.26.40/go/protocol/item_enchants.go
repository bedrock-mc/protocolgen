// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemEnchants struct {
	Slot         int32
	ItemEnchants [3][]EnchantmentInstance
}

// Marshal reads or writes ItemEnchants using its canonical wire layout.
func (x *ItemEnchants) Marshal(io IO) {
	io.Int32(&x.Slot)
	for index1 := range x.ItemEnchants {
		Slice(io, &x.ItemEnchants[index1])
	}
}
