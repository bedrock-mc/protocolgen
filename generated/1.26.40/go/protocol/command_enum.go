// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandEnum struct {
	Name   string
	Values []uint32
}

// Marshal reads or writes CommandEnum using its canonical wire layout.
func (x *CommandEnum) Marshal(io IO) {
	io.String(&x.Name)
	FuncSlice(io, &x.Values, io.Varuint32, io.Uint32)
}
