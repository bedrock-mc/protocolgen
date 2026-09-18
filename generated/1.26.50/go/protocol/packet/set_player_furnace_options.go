// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type SetPlayerFurnaceOptions struct {
	FurnaceType    protocol.FurnaceType
	FurnaceOptions protocol.FurnaceOptions
}

// ID returns the protocol ID for SetPlayerFurnaceOptions.
func (*SetPlayerFurnaceOptions) ID() uint32 { return IDSetPlayerFurnaceOptions }

// Marshal reads or writes SetPlayerFurnaceOptions using its canonical wire layout.
func (pk *SetPlayerFurnaceOptions) Marshal(io protocol.IO) {
	pk.FurnaceType.Marshal(io)
	pk.FurnaceOptions.Marshal(io)
}
