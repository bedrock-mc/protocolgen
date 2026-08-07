// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ContainerClose struct {
	ContainerID          uint8
	ContainerType        uint8
	ServerInitiatedClose bool
}

// Marshal reads or writes ContainerClose using its canonical wire layout.
func (x *ContainerClose) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Uint8(&x.ContainerType)
	io.Bool(&x.ServerInitiatedClose)
}

// ID returns the protocol ID for ContainerClose.
func (*ContainerClose) ID() uint32 { return IDContainerClose }
