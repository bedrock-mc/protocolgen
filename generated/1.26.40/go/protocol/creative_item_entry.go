// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CreativeItemEntry struct {
	CreativeNetID CreativeItemNetID
	ItemInstance  NetworkItemInstanceDescriptorSerializedData
	GroupIndex    uint32
}

// Marshal reads or writes CreativeItemEntry using its canonical wire layout.
func (x *CreativeItemEntry) Marshal(io IO) {
	x.CreativeNetID.Marshal(io)
	x.ItemInstance.Marshal(io)
	io.Varuint32(&x.GroupIndex)
}
