// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ItemRegistry struct {
	ItemData []protocol.ItemData
}

// Marshal reads or writes ItemRegistry using its canonical wire layout.
func (x *ItemRegistry) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.ItemData)
}

// ID returns the protocol ID for ItemRegistry.
func (*ItemRegistry) ID() uint32 { return IDItemRegistry }
