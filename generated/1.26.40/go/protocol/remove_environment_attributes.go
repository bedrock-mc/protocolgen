// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RemoveEnvironmentAttributes struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []string
}

func (*RemoveEnvironmentAttributes) isAttributeLayerSyncData() {}

// Marshal reads or writes RemoveEnvironmentAttributes using its canonical wire layout.
func (x *RemoveEnvironmentAttributes) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	FuncSlice(io, &x.Attributes, io.Varuint32, io.String)
}
