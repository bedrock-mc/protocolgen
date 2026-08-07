// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackResponseContainerInfo struct {
	FullContainerName FullContainerName
	Slots             []ItemStackResponseSlotInfo
}

// Marshal reads or writes ItemStackResponseContainerInfo using its canonical wire layout.
func (x *ItemStackResponseContainerInfo) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	FuncSlice(io, &x.Slots, io.Varuint32, func(value *ItemStackResponseSlotInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
