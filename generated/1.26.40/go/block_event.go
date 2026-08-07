// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BlockEvent struct {
	BlockPosition BlockPos
	EventType     int32
	EventValue    int32
}

// Marshal reads or writes BlockEvent using its canonical wire layout.
func (x *BlockEvent) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.Varint32(&x.EventType)
	io.Varint32(&x.EventValue)
}
