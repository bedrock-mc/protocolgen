// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type FeatureRegistryFeatureBinaryJsonFormat struct {
	FeatureName      string
	BinaryJsonOutput string
}

// Marshal reads or writes FeatureRegistryFeatureBinaryJsonFormat using its canonical wire layout.
func (x *FeatureRegistryFeatureBinaryJsonFormat) Marshal(io IO) {
	io.String(&x.FeatureName)
	io.String(&x.BinaryJsonOutput)
}
