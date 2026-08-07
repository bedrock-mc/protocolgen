// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ContainerRegistryCleanup is sent by the server to trigger a client-side cleanup of the dynamic
// container registry.
type ContainerRegistryCleanup struct {
	// RemovedContainers is a list of protocol.FullContainerName's that should be removed from the
	// client-side container registry.
	RemovedContainers []protocol.FullContainerName
}

// Marshal reads or writes ContainerRegistryCleanup using its canonical wire layout.
func (x *ContainerRegistryCleanup) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.RemovedContainers)
}

// ID returns the protocol ID for ContainerRegistryCleanup.
func (*ContainerRegistryCleanup) ID() uint32 { return IDContainerRegistryCleanup }
