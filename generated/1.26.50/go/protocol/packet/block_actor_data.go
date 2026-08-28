// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// BlockActorData is sent by the server to update data of a block entity client-side, for example
// the data of a chest.
type BlockActorData struct {
	BlockPosition protocol.BlockPos
	ActorDataTags []byte
}

// Marshal reads or writes BlockActorData using its canonical wire layout.
func (x *BlockActorData) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
	io.NBT(&x.ActorDataTags, protocol.NBTNetwork)
}

// ID returns the protocol ID for BlockActorData.
func (*BlockActorData) ID() uint32 { return IDBlockActorData }
