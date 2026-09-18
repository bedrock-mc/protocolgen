// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// SetPlayerInventoryOptions is a bidirectional packet that can be used to update the inventory options of a
// player.
type SetPlayerInventoryOptions struct {
	LeftInventoryTab  protocol.InventoryLeftTabIndex
	RightInventoryTab protocol.InventoryRightTabIndex
	Filtering         bool
	LayoutInv         protocol.InventoryLayout
	LayoutCraft       protocol.InventoryLayout
}

// ID returns the protocol ID for SetPlayerInventoryOptions.
func (*SetPlayerInventoryOptions) ID() uint32 { return IDSetPlayerInventoryOptions }

// Marshal reads or writes SetPlayerInventoryOptions using its canonical wire layout.
func (pk *SetPlayerInventoryOptions) Marshal(io protocol.IO) {
	pk.LeftInventoryTab.Marshal(io)
	pk.RightInventoryTab.Marshal(io)
	io.Bool(&pk.Filtering)
	pk.LayoutInv.Marshal(io)
	pk.LayoutCraft.Marshal(io)
}
