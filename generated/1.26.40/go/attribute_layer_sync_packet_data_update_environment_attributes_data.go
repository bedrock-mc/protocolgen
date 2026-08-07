// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) isAttributeLayerSyncPacketData() {}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	FuncSlice(io, &x.Attributes, io.Varuint32, func(value *EASEnvironmentAttributeData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
