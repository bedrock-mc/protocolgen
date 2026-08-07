// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ContainerMixDataEntry struct {
	FromItemID    int32
	ReagentItemID int32
	ToItemID      int32
}

// Marshal reads or writes ContainerMixDataEntry using its canonical wire layout.
func (x *ContainerMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemID)
	io.Varint32(&x.ReagentItemID)
	io.Varint32(&x.ToItemID)
}
