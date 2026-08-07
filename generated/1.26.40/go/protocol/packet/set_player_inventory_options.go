// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetPlayerInventoryOptions struct {
	InventoryOptions protocol.InventoryOptions
}

// Marshal reads or writes SetPlayerInventoryOptions using its canonical wire layout.
func (x *SetPlayerInventoryOptions) Marshal(io protocol.IO) {
	x.InventoryOptions.Marshal(io)
}

// ID returns the protocol ID for SetPlayerInventoryOptions.
func (*SetPlayerInventoryOptions) ID() uint32 { return IDSetPlayerInventoryOptions }
