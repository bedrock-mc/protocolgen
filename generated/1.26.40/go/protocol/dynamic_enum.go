// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicEnum struct {
	EnumName    string
	EnumOptions []string
}

// Marshal reads or writes DynamicEnum using its canonical wire layout.
func (x *DynamicEnum) Marshal(io IO) {
	io.String(&x.EnumName)
	FuncSlice(io, &x.EnumOptions, io.Varuint32, io.String)
}
