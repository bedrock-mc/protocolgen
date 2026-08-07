// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeTagsData struct {
	Tags []uint16
}

// Marshal reads or writes BiomeTagsData using its canonical wire layout.
func (x *BiomeTagsData) Marshal(io IO) {
	FuncSlice(io, &x.Tags, io.Varuint32, io.Uint16)
}
