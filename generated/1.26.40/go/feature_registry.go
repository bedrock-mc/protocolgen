// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type FeatureRegistry struct {
	FeaturesDataList []FeatureRegistryFeatureBinaryJsonFormat
}

// Marshal reads or writes FeatureRegistry using its canonical wire layout.
func (x *FeatureRegistry) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.FeaturesDataList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.FeaturesDataList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.FeaturesDataList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.FeaturesDataList = make([]FeatureRegistryFeatureBinaryJsonFormat, int(count1))
	}
	for index2 := range x.FeaturesDataList {
		x.FeaturesDataList[index2].Marshal(io)
	}
}
