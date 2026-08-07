// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerEnchantOptions struct {
	Options []ItemEnchantOption
}

// Marshal reads or writes PlayerEnchantOptions using its canonical wire layout.
func (x *PlayerEnchantOptions) Marshal(io IO) {
	FuncSlice(io, &x.Options, io.Varuint32, func(value *ItemEnchantOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
