// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []string
}

func (AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData) isAttributeLayerSyncPacketData() {}

// Marshal reads or writes AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	FuncSlice(io, &x.Attributes, io.Varuint32, io.String)
}
