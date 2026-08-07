// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableCommands struct {
	EnumValues              []string
	ChainedSubcommandValues []string
	PostFixes               []string
	EnumData                []AvailableCommandsEnumData
	ChainedSubcommandData   []AvailableCommandsChainedSubcommandData
	Commands                []AvailableCommandsPacketCommandData
	SoftEnums               []AvailableCommandsSoftEnumData
	Constraints             []AvailableCommandsConstrainedValueData
}

// Marshal reads or writes AvailableCommands using its canonical wire layout.
func (x *AvailableCommands) Marshal(io IO) {
	FuncSlice(io, &x.EnumValues, io.Varuint32, io.String)
	FuncSlice(io, &x.ChainedSubcommandValues, io.Varuint32, io.String)
	FuncSlice(io, &x.PostFixes, io.Varuint32, io.String)
	FuncSlice(io, &x.EnumData, io.Varuint32, func(value *AvailableCommandsEnumData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ChainedSubcommandData, io.Varuint32, func(value *AvailableCommandsChainedSubcommandData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.Commands, io.Varuint32, func(value *AvailableCommandsPacketCommandData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.SoftEnums, io.Varuint32, func(value *AvailableCommandsSoftEnumData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.Constraints, io.Varuint32, func(value *AvailableCommandsConstrainedValueData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
