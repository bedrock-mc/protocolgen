// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type BlockEvent struct {
	BlockPosition protocol.BlockPos
	EventType     int32
	EventValue    int32
}

// Marshal reads or writes BlockEvent using its canonical wire layout.
func (x *BlockEvent) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
	io.Varint32(&x.EventType)
	io.Varint32(&x.EventValue)
}
