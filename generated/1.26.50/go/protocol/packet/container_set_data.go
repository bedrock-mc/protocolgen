// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// ContainerSetData is sent by the server to update specific data of a single container, meaning a
// block such as a furnace or a brewing stand. This data is usually used by the client to display
// certain features client-side.
type ContainerSetData struct {
	ContainerID uint8
	IDValue     int32
	// Value is the value of the property. Its use differs per property.
	Value int32
}

// Marshal reads or writes ContainerSetData using its canonical wire layout.
func (x *ContainerSetData) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	protocol.Minimum(io, &x.ContainerID, 0)
	protocol.Maximum(io, &x.ContainerID, 255)
	io.Varint32(&x.IDValue)
	io.Varint32(&x.Value)
}

// ID returns the protocol ID for ContainerSetData.
func (*ContainerSetData) ID() uint32 { return IDContainerSetData }
