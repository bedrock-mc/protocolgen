// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackResponseContainerInfo struct {
	FullContainerName FullContainerName
	Slots             []ItemStackResponseSlotInfo
}

// Marshal reads or writes ItemStackResponseContainerInfo using its canonical wire layout.
func (x *ItemStackResponseContainerInfo) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	Slice(io, &x.Slots)
}
