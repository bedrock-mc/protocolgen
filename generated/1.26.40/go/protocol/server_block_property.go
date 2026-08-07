// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerBlockProperty struct {
	BlockName       string
	BlockDefinition []byte
}

// Marshal reads or writes ServerBlockProperty using its canonical wire layout.
func (x *ServerBlockProperty) Marshal(io IO) {
	io.String(&x.BlockName)
	io.NBT(&x.BlockDefinition, NBTNetwork)
}
