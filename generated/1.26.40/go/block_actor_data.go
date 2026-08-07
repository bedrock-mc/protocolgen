// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BlockActorData struct {
	BlockPosition BlockPos
	ActorDataTags []byte
}

// Marshal reads or writes BlockActorData using its canonical wire layout.
func (x *BlockActorData) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.NBT(&x.ActorDataTags)
}
