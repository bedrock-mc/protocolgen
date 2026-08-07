// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChainedSubcommand struct {
	Name             string
	SubCommandValues []ChainedSubcommandValue
}

// Marshal reads or writes ChainedSubcommand using its canonical wire layout.
func (x *ChainedSubcommand) Marshal(io IO) {
	io.String(&x.Name)
	Slice(io, &x.SubCommandValues)
}
