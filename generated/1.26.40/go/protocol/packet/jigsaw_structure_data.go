// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type JigsawStructureData struct {
	JigsawStructureDataTag []byte
}

// Marshal reads or writes JigsawStructureData using its canonical wire layout.
func (x *JigsawStructureData) Marshal(io protocol.IO) {
	io.NBT(&x.JigsawStructureDataTag, protocol.NBTNetwork)
}

// ID returns the protocol ID for JigsawStructureData.
func (*JigsawStructureData) ID() uint32 { return IDJigsawStructureData }
