// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AttributeLayerSyncPacketDataUpdateAttributeLayersData struct {
	AttributeLayers []EASAttributeLayerData
}

func (AttributeLayerSyncPacketDataUpdateAttributeLayersData) isAttributeLayerSyncPacketData() {}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateAttributeLayersData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateAttributeLayersData) Marshal(io IO) {
	FuncSlice(io, &x.AttributeLayers, io.Varuint32, func(value *EASAttributeLayerData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
