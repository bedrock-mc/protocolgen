// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundAttributeLayerSync struct {
	Data protocol.AttributeLayerSyncPacketData
}

// Marshal reads or writes ClientboundAttributeLayerSync using its canonical wire layout.
func (x *ClientboundAttributeLayerSync) Marshal(io protocol.IO) {
	protocol.MarshalAttributeLayerSyncPacketData(io, &x.Data)
}
