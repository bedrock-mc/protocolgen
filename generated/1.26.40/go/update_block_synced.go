// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateBlockSynced struct {
	BlockPosition    BlockPos
	BlockRuntimeID   uint32
	Flags            uint32
	Layer            uint32
	UniqueActorId    uint64
	ActorSyncMessage uint64
}

// Marshal reads or writes UpdateBlockSynced using its canonical wire layout.
func (x *UpdateBlockSynced) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.Varuint32(&x.BlockRuntimeID)
	io.Varuint32(&x.Flags)
	io.Varuint32(&x.Layer)
	io.Varuint64(&x.UniqueActorId)
	io.Varuint64(&x.ActorSyncMessage)
}
