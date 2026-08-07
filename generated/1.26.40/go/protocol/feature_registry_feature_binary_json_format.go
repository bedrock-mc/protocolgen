// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type FeatureRegistryFeatureBinaryJSONFormat struct {
	FeatureName      string
	BinaryJSONOutput string
}

// Marshal reads or writes FeatureRegistryFeatureBinaryJSONFormat using its canonical wire layout.
func (x *FeatureRegistryFeatureBinaryJSONFormat) Marshal(io IO) {
	io.String(&x.FeatureName)
	io.String(&x.BinaryJSONOutput)
}
