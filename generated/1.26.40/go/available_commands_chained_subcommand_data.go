// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableCommandsChainedSubcommandData struct {
	Name             string
	SubCommandValues []AvailableCommandsChainedSubcommandRelationship
}

// Marshal reads or writes AvailableCommandsChainedSubcommandData using its canonical wire layout.
func (x *AvailableCommandsChainedSubcommandData) Marshal(io IO) {
	io.String(&x.Name)
	FuncSlice(io, &x.SubCommandValues, io.Varuint32, func(value *AvailableCommandsChainedSubcommandRelationship) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
