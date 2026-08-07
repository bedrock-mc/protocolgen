// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GuiDataPickItem struct {
	ItemName       string
	ItemEffectName string
	Slot           int32
}

// Marshal reads or writes GuiDataPickItem using its canonical wire layout.
func (x *GuiDataPickItem) Marshal(io IO) {
	io.String(&x.ItemName)
	io.String(&x.ItemEffectName)
	io.Int32(&x.Slot)
}
