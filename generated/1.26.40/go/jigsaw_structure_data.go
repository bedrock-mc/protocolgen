// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type JigsawStructureData struct {
	JigsawStructureDataTag []byte
}

// Marshal reads or writes JigsawStructureData using its canonical wire layout.
func (x *JigsawStructureData) Marshal(io IO) {
	io.NBT(&x.JigsawStructureDataTag)
}
