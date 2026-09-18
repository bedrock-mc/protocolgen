// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ClientboundAttributeLayerSync struct {
	Data protocol.AttributeLayerSyncData
}

// ID ...
func (*ClientboundAttributeLayerSync) ID() uint32 {
	return IDClientboundAttributeLayerSync
}

func (pk *ClientboundAttributeLayerSync) Marshal(io protocol.IO) {
	protocol.MarshalAttributeLayerSyncData(io, &pk.Data)
}
