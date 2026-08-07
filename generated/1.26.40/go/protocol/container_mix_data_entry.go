// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ContainerMixDataEntry struct {
	FromItemId    int32
	ReagentItemId int32
	ToItemId      int32
}

// Marshal reads or writes ContainerMixDataEntry using its canonical wire layout.
func (x *ContainerMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemId)
	io.Varint32(&x.ReagentItemId)
	io.Varint32(&x.ToItemId)
}
