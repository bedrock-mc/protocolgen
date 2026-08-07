// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type FeatureRegistry struct {
	FeaturesDataList []protocol.FeatureRegistryFeatureBinaryJsonFormat
}

// Marshal reads or writes FeatureRegistry using its canonical wire layout.
func (x *FeatureRegistry) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.FeaturesDataList, io.Varuint32, func(value *protocol.FeatureRegistryFeatureBinaryJsonFormat) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for FeatureRegistry.
func (*FeatureRegistry) ID() uint32 { return IDFeatureRegistry }
