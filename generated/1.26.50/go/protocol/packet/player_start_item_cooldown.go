// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type PlayerStartItemCooldown struct {
	ItemCategory  string
	DurationTicks int32
}

// ID returns the protocol ID for PlayerStartItemCooldown.
func (*PlayerStartItemCooldown) ID() uint32 { return IDPlayerStartItemCooldown }

// Marshal reads or writes PlayerStartItemCooldown using its canonical wire layout.
func (pk *PlayerStartItemCooldown) Marshal(io protocol.IO) {
	io.String(&pk.ItemCategory)
	io.Varint32(&pk.DurationTicks)
}
