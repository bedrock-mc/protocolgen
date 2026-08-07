// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerEnchantOptions struct {
	Options []protocol.ItemEnchantOption
}

// Marshal reads or writes PlayerEnchantOptions using its canonical wire layout.
func (x *PlayerEnchantOptions) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.Options, io.Varuint32, func(value *protocol.ItemEnchantOption) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for PlayerEnchantOptions.
func (*PlayerEnchantOptions) ID() uint32 { return IDPlayerEnchantOptions }
