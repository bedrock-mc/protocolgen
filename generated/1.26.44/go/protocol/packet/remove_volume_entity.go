// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// RemoveVolumeEntity indicates a volume entity to be removed from server to client.
type RemoveVolumeEntity struct {
	EntityNetworkID protocol.EntityNetID
	DimensionType   protocol.DimensionType
}

// ID returns the protocol ID for RemoveVolumeEntity.
func (*RemoveVolumeEntity) ID() uint32 { return IDRemoveVolumeEntity }

// Marshal reads or writes RemoveVolumeEntity using its canonical wire layout.
func (pk *RemoveVolumeEntity) Marshal(io protocol.IO) {
	pk.EntityNetworkID.Marshal(io)
	pk.DimensionType.Marshal(io)
}
