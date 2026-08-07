// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CreativeGroupInfo struct {
	CreativeCategory CreativeItemCategory
	Name             string
	GroupIconItem    CerealizerNetworkItemInstanceDescriptorSerializedData
}

// Marshal reads or writes CreativeGroupInfo using its canonical wire layout.
func (x *CreativeGroupInfo) Marshal(io IO) {
	IntegerFunc(&x.CreativeCategory, io.Uint8)
	io.String(&x.Name)
	x.GroupIconItem.Marshal(io)
}
