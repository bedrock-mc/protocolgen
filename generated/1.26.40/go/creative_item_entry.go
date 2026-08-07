// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CreativeItemEntry struct {
	CreativeNetId TypedServerNetIdStructCreativeItemNetIdTag
	ItemInstance  CerealizerNetworkItemInstanceDescriptorSerializedData
	GroupIndex    uint32
}

// Marshal reads or writes CreativeItemEntry using its canonical wire layout.
func (x *CreativeItemEntry) Marshal(io IO) {
	x.CreativeNetId.Marshal(io)
	x.ItemInstance.Marshal(io)
	io.Varuint32(&x.GroupIndex)
}
