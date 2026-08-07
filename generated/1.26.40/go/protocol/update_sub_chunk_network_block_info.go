// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type UpdateSubChunkNetworkBlockInfo struct {
	Pos                       BlockPos
	RuntimeId                 uint32
	UpdateFlags               uint32
	SyncMessageEntityUniqueID uint64
	SyncMessageMessage        uint32
}

// Marshal reads or writes UpdateSubChunkNetworkBlockInfo using its canonical wire layout.
func (x *UpdateSubChunkNetworkBlockInfo) Marshal(io IO) {
	x.Pos.Marshal(io)
	io.Varuint32(&x.RuntimeId)
	io.Varuint32(&x.UpdateFlags)
	io.Varuint64(&x.SyncMessageEntityUniqueID)
	io.Varuint32(&x.SyncMessageMessage)
}
