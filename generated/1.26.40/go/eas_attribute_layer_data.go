// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EASAttributeLayerData struct {
	Name       string
	NoiseName  Optional[string]
	Dimension  DimensionType
	Settings   EASAttributeLayerSettings
	Attributes []EASEnvironmentAttributeData
}

// Marshal reads or writes EASAttributeLayerData using its canonical wire layout.
func (x *EASAttributeLayerData) Marshal(io IO) {
	io.String(&x.Name)
	OptionalFunc(io, &x.NoiseName, io.String)
	x.Dimension.Marshal(io)
	x.Settings.Marshal(io)
	FuncSlice(io, &x.Attributes, io.Varuint32, func(value *EASEnvironmentAttributeData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
