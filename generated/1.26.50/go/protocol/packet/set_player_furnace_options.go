// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type SetPlayerFurnaceOptions struct {
	FurnaceType    protocol.FurnaceType
	FurnaceOptions protocol.FurnaceOptions
}

// ID ...
func (*SetPlayerFurnaceOptions) ID() uint32 {
	return IDSetPlayerFurnaceOptions
}

func (pk *SetPlayerFurnaceOptions) Marshal(io protocol.IO) {
	pk.FurnaceType.Marshal(io)
	pk.FurnaceOptions.Marshal(io)
}
