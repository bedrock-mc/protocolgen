// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableCommandsOverloadData struct {
	IsChaining    bool
	ParameterData []AvailableCommandsParamData
}

// Marshal reads or writes AvailableCommandsOverloadData using its canonical wire layout.
func (x *AvailableCommandsOverloadData) Marshal(io IO) {
	io.Bool(&x.IsChaining)
	FuncSlice(io, &x.ParameterData, io.Varuint32, func(value *AvailableCommandsParamData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
