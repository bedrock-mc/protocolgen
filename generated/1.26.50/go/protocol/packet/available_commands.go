// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// AvailableCommands is sent by the server to send a list of all commands that the player is able to
// use on the server. This packet holds all the arguments of each commands as well, making it
// possible for the client to provide auto-completion and command usages. AvailableCommands packets
// can be resent, but the packet is often very big, so doing this very often should be avoided.
type AvailableCommands struct {
	// EnumValues is a slice of all enum values of any enum in the AvailableCommands packet. EnumValues
	// generally should contain each possible value only once. Enums are built by pointing to entries in
	// this slice.
	EnumValues []string
	// ChainedSubcommandValues is a slice of all chained subcommand names. ChainedSubcommandValues
	// generally should contain each possible value only once. ChainedSubcommands are built by pointing
	// to entries in this slice.
	ChainedSubcommandValues []string
	// PostFixes, like EnumValues, is a slice of all suffix values of any command parameter in the
	// AvailableCommands packet.
	PostFixes []string
	// EnumData is a slice of all (fixed) command enums present in any of the commands.
	EnumData []protocol.CommandEnum
	// ChainedSubcommandData is a slice of all subcommands that are followed by a chained command. An
	// example usage of this is /execute which allows you to run another command as another entity or at
	// a different position etc.
	ChainedSubcommandData []protocol.ChainedSubcommand
	// Commands is a list of all commands that the client should show client-side. The AvailableCommands
	// packet replaces any commands sent before. It does not only add the commands that are sent in it.
	Commands []protocol.Command
	// SoftEnums is a slice of dynamic command enums. These command enums can be changed during runtime
	// without having to resend an AvailableCommands packet.
	SoftEnums []protocol.DynamicEnum
	// Constraints is a list of constraints that should be applied to certain options of enums in the
	// commands above.
	Constraints []protocol.CommandEnumConstraint
}

// Marshal reads or writes AvailableCommands using its canonical wire layout.
func (x *AvailableCommands) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.EnumValues, io.Varuint32, io.String)
	protocol.FuncSlice(io, &x.ChainedSubcommandValues, io.Varuint32, io.String)
	protocol.FuncSlice(io, &x.PostFixes, io.Varuint32, io.String)
	protocol.Slice(io, &x.EnumData)
	protocol.SliceLimits(io, &x.ChainedSubcommandData, 0, 16)
	protocol.Slice(io, &x.Commands)
	protocol.Slice(io, &x.SoftEnums)
	protocol.Slice(io, &x.Constraints)
}

// ID returns the protocol ID for AvailableCommands.
func (*AvailableCommands) ID() uint32 { return IDAvailableCommands }
