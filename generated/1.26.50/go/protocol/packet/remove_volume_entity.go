// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// RemoveVolumeEntity indicates a volume entity to be removed from server to client.
type RemoveVolumeEntity struct {
	EntityNetworkID protocol.EntityNetID
	DimensionType   protocol.DimensionType
}

// Marshal reads or writes RemoveVolumeEntity using its canonical wire layout.
func (x *RemoveVolumeEntity) Marshal(io protocol.IO) {
	x.EntityNetworkID.Marshal(io)
	x.DimensionType.Marshal(io)
}

// ID returns the protocol ID for RemoveVolumeEntity.
func (*RemoveVolumeEntity) ID() uint32 { return IDRemoveVolumeEntity }
