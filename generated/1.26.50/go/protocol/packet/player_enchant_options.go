// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// PlayerEnchantOptions is sent by the server to update the enchantment options displayed when the
// user opens the enchantment table and puts an item in. This packet was added in 1.16 and allows
// the server to decide on the enchantments that can be selected by the player. The
// PlayerEnchantOptions packet should be sent once for every slot update of the enchantment table.
// The vanilla server sends an empty PlayerEnchantOptions packet when the player opens the
// enchantment table (air is present in the enchantment table slot) and sends the packet with actual
// enchantments in it when items are put in that can have enchantments.
type PlayerEnchantOptions struct {
	// Options is a list of possible enchantment options for the item that was put into the enchantment
	// table.
	Options []protocol.ItemEnchantOption
}

// Marshal reads or writes PlayerEnchantOptions using its canonical wire layout.
func (x *PlayerEnchantOptions) Marshal(io protocol.IO) {
	protocol.SliceLimits(io, &x.Options, 0, 3)
}

// ID returns the protocol ID for PlayerEnchantOptions.
func (*PlayerEnchantOptions) ID() uint32 { return IDPlayerEnchantOptions }
