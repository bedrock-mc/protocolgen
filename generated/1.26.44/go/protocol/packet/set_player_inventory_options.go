// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// SetPlayerInventoryOptions is a bidirectional packet that can be used to update the inventory
// options of a player.
type SetPlayerInventoryOptions struct {
	InventoryOptions protocol.InventoryOptions
}

// Marshal reads or writes SetPlayerInventoryOptions using its canonical wire layout.
func (x *SetPlayerInventoryOptions) Marshal(io protocol.IO) {
	x.InventoryOptions.Marshal(io)
}

// ID returns the protocol ID for SetPlayerInventoryOptions.
func (*SetPlayerInventoryOptions) ID() uint32 { return IDSetPlayerInventoryOptions }
