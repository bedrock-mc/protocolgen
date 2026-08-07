// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeStringList struct {
	Strings []string
}

// Marshal reads or writes BiomeStringList using its canonical wire layout.
func (x *BiomeStringList) Marshal(io IO) {
	FuncSlice(io, &x.Strings, io.Varuint32, io.String)
}
