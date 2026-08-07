// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AvailableCommands struct {
	EnumValues              []string
	ChainedSubcommandValues []string
	PostFixes               []string
	EnumData                []protocol.AvailableCommandsEnumData
	ChainedSubcommandData   []protocol.AvailableCommandsChainedSubcommandData
	Commands                []protocol.AvailableCommandsPacketCommandData
	SoftEnums               []protocol.AvailableCommandsSoftEnumData
	Constraints             []protocol.AvailableCommandsConstrainedValueData
}

// Marshal reads or writes AvailableCommands using its canonical wire layout.
func (x *AvailableCommands) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.EnumValues, io.Varuint32, io.String)
	protocol.FuncSlice(io, &x.ChainedSubcommandValues, io.Varuint32, io.String)
	protocol.FuncSlice(io, &x.PostFixes, io.Varuint32, io.String)
	protocol.Slice(io, &x.EnumData)
	protocol.Slice(io, &x.ChainedSubcommandData)
	protocol.Slice(io, &x.Commands)
	protocol.Slice(io, &x.SoftEnums)
	protocol.Slice(io, &x.Constraints)
}

// ID returns the protocol ID for AvailableCommands.
func (*AvailableCommands) ID() uint32 { return IDAvailableCommands }
