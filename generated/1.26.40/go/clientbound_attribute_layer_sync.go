// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundAttributeLayerSync struct {
	Data AttributeLayerSyncPacketData
}

// Marshal reads or writes ClientboundAttributeLayerSync using its canonical wire layout.
func (x *ClientboundAttributeLayerSync) Marshal(io IO) {
	marshalAttributeLayerSyncPacketData(io, &x.Data)
}
