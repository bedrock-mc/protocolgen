// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// FeatureRegistry is a packet used to notify the client about the world generation features the
// server is currently using. This is used in combination with the client-side world generation
// system introduced in v1.19.20, allowing the client to completely generate the chunks of the world
// without having to rely on the server.
type FeatureRegistry struct {
	// FeaturesDataList is a slice of all registered world generation features.
	FeaturesDataList []protocol.FeatureRegistryFeatureBinaryJSONFormat
}

// Marshal reads or writes FeatureRegistry using its canonical wire layout.
func (x *FeatureRegistry) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.FeaturesDataList)
}

// ID returns the protocol ID for FeatureRegistry.
func (*FeatureRegistry) ID() uint32 { return IDFeatureRegistry }
