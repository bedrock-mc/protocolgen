// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (*AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) isAttributeLayerSyncPacketData() {
}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	Slice(io, &x.Attributes)
}
