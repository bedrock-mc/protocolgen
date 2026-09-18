// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// BlockActorData is sent by the server to update data of a block entity client-side, for example the data of
// a chest.
type BlockActorData struct {
	BlockPosition protocol.BlockPos
	ActorDataTags []byte
}

// ID ...
func (*BlockActorData) ID() uint32 {
	return IDBlockActorData
}

func (pk *BlockActorData) Marshal(io protocol.IO) {
	pk.BlockPosition.Marshal(io)
	io.NBT(&pk.ActorDataTags, protocol.NBTNetwork)
}
