// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerStartItemCooldown struct {
	ItemCategory  string
	DurationTicks int32
}

// Marshal reads or writes PlayerStartItemCooldown using its canonical wire layout.
func (x *PlayerStartItemCooldown) Marshal(io protocol.IO) {
	io.String(&x.ItemCategory)
	io.Varint32(&x.DurationTicks)
}

// ID returns the protocol ID for PlayerStartItemCooldown.
func (*PlayerStartItemCooldown) ID() uint32 { return IDPlayerStartItemCooldown }
