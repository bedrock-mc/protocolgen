// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerSyncPacketDataUpdateAttributeLayersData struct {
	AttributeLayers []EASAttributeLayerData
}

func (*AttributeLayerSyncPacketDataUpdateAttributeLayersData) isAttributeLayerSyncPacketData() {}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateAttributeLayersData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateAttributeLayersData) Marshal(io IO) {
	Slice(io, &x.AttributeLayers)
}
