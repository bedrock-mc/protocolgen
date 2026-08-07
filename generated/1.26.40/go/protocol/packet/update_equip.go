// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateEquip struct {
	ContainerID    uint8
	Type           uint8
	Size           int32
	EntityUniqueID int64
	Data           []byte
}

// Marshal reads or writes UpdateEquip using its canonical wire layout.
func (x *UpdateEquip) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Uint8(&x.Type)
	io.Varint32(&x.Size)
	io.ActorUniqueID(&x.EntityUniqueID)
	io.NBT(&x.Data)
}

// ID returns the protocol ID for UpdateEquip.
func (*UpdateEquip) ID() uint32 { return IDUpdateEquip }
