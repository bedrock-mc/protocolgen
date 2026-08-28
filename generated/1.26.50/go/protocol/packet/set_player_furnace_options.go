// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

type SetPlayerFurnaceOptions struct {
	FurnaceType    protocol.FurnaceType
	FurnaceOptions protocol.FurnaceOptions
}

// Marshal reads or writes SetPlayerFurnaceOptions using its canonical wire layout.
func (x *SetPlayerFurnaceOptions) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.FurnaceType, io.Uint8)
	x.FurnaceOptions.Marshal(io)
}

// ID returns the protocol ID for SetPlayerFurnaceOptions.
func (*SetPlayerFurnaceOptions) ID() uint32 { return IDSetPlayerFurnaceOptions }
