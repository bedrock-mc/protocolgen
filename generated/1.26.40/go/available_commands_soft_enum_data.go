// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableCommandsSoftEnumData struct {
	EnumName    string
	EnumOptions []string
}

// Marshal reads or writes AvailableCommandsSoftEnumData using its canonical wire layout.
func (x *AvailableCommandsSoftEnumData) Marshal(io IO) {
	io.String(&x.EnumName)
	FuncSlice(io, &x.EnumOptions, io.Varuint32, io.String)
}
