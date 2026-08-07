// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemRegistry struct {
	ItemData []ItemData
}

// Marshal reads or writes ItemRegistry using its canonical wire layout.
func (x *ItemRegistry) Marshal(io IO) {
	FuncSlice(io, &x.ItemData, io.Varuint32, func(value *ItemData) {
		value.Marshal(io)
	})
}
