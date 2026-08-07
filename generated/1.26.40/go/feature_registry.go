// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type FeatureRegistry struct {
	FeaturesDataList []FeatureRegistryFeatureBinaryJsonFormat
}

// Marshal reads or writes FeatureRegistry using its canonical wire layout.
func (x *FeatureRegistry) Marshal(io IO) {
	FuncSlice(io, &x.FeaturesDataList, io.Varuint32, func(value *FeatureRegistryFeatureBinaryJsonFormat) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
