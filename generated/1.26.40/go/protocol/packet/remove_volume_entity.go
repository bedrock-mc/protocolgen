// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RemoveVolumeEntity struct {
	EntityNetworkId protocol.EntityNetId
	DimensionType   protocol.DimensionType
}

// Marshal reads or writes RemoveVolumeEntity using its canonical wire layout.
func (x *RemoveVolumeEntity) Marshal(io protocol.IO) {
	x.EntityNetworkId.Marshal(io)
	x.DimensionType.Marshal(io)
}
