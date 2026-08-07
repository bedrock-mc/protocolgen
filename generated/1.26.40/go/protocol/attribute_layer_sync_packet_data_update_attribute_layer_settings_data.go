// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	AttributesLayerSettings EASAttributeLayerSettings
}

func (AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData) isAttributeLayerSyncPacketData() {
}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	x.AttributesLayerSettings.Marshal(io)
}
