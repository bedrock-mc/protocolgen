// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateSoftEnum struct {
	EnumName   string
	Values     []string
	UpdateType SoftEnumUpdateType
}

// Marshal reads or writes UpdateSoftEnum using its canonical wire layout.
func (x *UpdateSoftEnum) Marshal(io IO) {
	io.String(&x.EnumName)
	FuncSlice(io, &x.Values, io.Varuint32, io.String)
	IntegerFunc(&x.UpdateType, io.Uint8)
}
