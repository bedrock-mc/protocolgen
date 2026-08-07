// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ContainerSetData struct {
	ContainerID uint8
	IDValue     int32
	Value       int32
}

// Marshal reads or writes ContainerSetData using its canonical wire layout.
func (x *ContainerSetData) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Varint32(&x.IDValue)
	io.Varint32(&x.Value)
}

// ID returns the protocol ID for ContainerSetData.
func (*ContainerSetData) ID() uint32 { return IDContainerSetData }
