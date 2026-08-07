// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsParamData struct {
	Name        string
	ParseSymbol uint32
	IsOptional  bool
	Options     uint8
}

// Marshal reads or writes AvailableCommandsParamData using its canonical wire layout.
func (x *AvailableCommandsParamData) Marshal(io IO) {
	io.String(&x.Name)
	io.Uint32(&x.ParseSymbol)
	io.Bool(&x.IsOptional)
	io.Uint8(&x.Options)
}
