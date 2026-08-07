// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerData struct {
	AttributeLayers []EASAttributeLayerData
}

func (*AttributeLayerData) isAttributeLayerSyncData() {}

// Marshal reads or writes AttributeLayerData using its canonical wire layout.
func (x *AttributeLayerData) Marshal(io IO) {
	Slice(io, &x.AttributeLayers)
}
