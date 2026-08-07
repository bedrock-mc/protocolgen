// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EnvironmentAttributeData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (*EnvironmentAttributeData) isAttributeLayerSyncData() {}

// Marshal reads or writes EnvironmentAttributeData using its canonical wire layout.
func (x *EnvironmentAttributeData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	Slice(io, &x.Attributes)
}
