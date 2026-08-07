// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandParameter struct {
	Name        string
	ParseSymbol uint32
	IsOptional  bool
	Options     uint8
}

// Marshal reads or writes CommandParameter using its canonical wire layout.
func (x *CommandParameter) Marshal(io IO) {
	io.String(&x.Name)
	io.Uint32(&x.ParseSymbol)
	io.Bool(&x.IsOptional)
	io.Uint8(&x.Options)
}
