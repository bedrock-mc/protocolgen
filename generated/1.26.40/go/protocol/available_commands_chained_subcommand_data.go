// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsChainedSubcommandData struct {
	Name             string
	SubCommandValues []AvailableCommandsChainedSubcommandRelationship
}

// Marshal reads or writes AvailableCommandsChainedSubcommandData using its canonical wire layout.
func (x *AvailableCommandsChainedSubcommandData) Marshal(io IO) {
	io.String(&x.Name)
	Slice(io, &x.SubCommandValues)
}
