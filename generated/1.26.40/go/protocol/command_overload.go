// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandOverload struct {
	IsChaining    bool
	ParameterData []CommandParameter
}

// Marshal reads or writes CommandOverload using its canonical wire layout.
func (x *CommandOverload) Marshal(io IO) {
	io.Bool(&x.IsChaining)
	Slice(io, &x.ParameterData)
}
