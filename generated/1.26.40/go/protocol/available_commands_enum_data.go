// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsEnumData struct {
	Name   string
	Values []uint32
}

// Marshal reads or writes AvailableCommandsEnumData using its canonical wire layout.
func (x *AvailableCommandsEnumData) Marshal(io IO) {
	io.String(&x.Name)
	FuncSlice(io, &x.Values, io.Varuint32, io.Uint32)
}
