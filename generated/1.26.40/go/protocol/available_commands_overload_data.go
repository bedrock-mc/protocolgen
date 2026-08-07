// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsOverloadData struct {
	IsChaining    bool
	ParameterData []AvailableCommandsParamData
}

// Marshal reads or writes AvailableCommandsOverloadData using its canonical wire layout.
func (x *AvailableCommandsOverloadData) Marshal(io IO) {
	io.Bool(&x.IsChaining)
	Slice(io, &x.ParameterData)
}
