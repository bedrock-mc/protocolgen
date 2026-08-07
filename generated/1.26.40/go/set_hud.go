// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetHud struct {
	HudElement []HudElement
	HudVisible HudVisibility
}

// Marshal reads or writes SetHud using its canonical wire layout.
func (x *SetHud) Marshal(io IO) {
	FuncSlice(io, &x.HudElement, io.Varuint32, func(value *HudElement) {
		item := *value
		IntegerFunc(&item, io.Varint32)
		*value = item
	})
	IntegerFunc(&x.HudVisible, io.Varint32)
}
