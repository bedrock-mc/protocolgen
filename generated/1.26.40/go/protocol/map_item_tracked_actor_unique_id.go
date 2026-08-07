// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MapItemTrackedActorUniqueId struct {
	Type          MapItemTrackedActorType
	EntityID      Optional[int64]
	BlockPosition Optional[BlockPos]
}

// Marshal reads or writes MapItemTrackedActorUniqueId using its canonical wire layout.
func (x *MapItemTrackedActorUniqueId) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Int32)
	OptionalFunc(io, &x.EntityID, io.ActorUniqueID)
	OptionalFunc(io, &x.BlockPosition, func(value *BlockPos) {
		value.Marshal(io)
	})
}
