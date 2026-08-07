// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

// ChainedSubcommand represents a subcommand that can have chained commands, such as /execute which
// allows you to run another command as another entity or at a different position etc.
type ChainedSubcommand struct {
	// Name is the name of the chained subcommand and shows up in the list as a regular subcommand enum.
	Name string
	// SubCommandValues contains the index and parameter type of the chained subcommand.
	SubCommandValues []ChainedSubcommandValue
}

// Marshal reads or writes ChainedSubcommand using its canonical wire layout.
func (x *ChainedSubcommand) Marshal(io IO) {
	io.String(&x.Name)
	Slice(io, &x.SubCommandValues)
}

// ChainedSubcommandValue represents the value for a chained subcommand argument.
type ChainedSubcommandValue struct {
	// SubCommandFirstValue is the index of the argument in the ChainedSubcommandValues slice from the
	// AvailableCommands packet. This is then used to set the type specified by the Value field below.
	SubCommandFirstValue uint32
	// SubCommandSecondValue is a combination of the flags above and specified the type of argument.
	// Unlike regular parameter types, this should NOT contain any of the special flags (valid, enum,
	// suffixed or soft enum) but only the basic types.
	SubCommandSecondValue uint32
}

// Marshal reads or writes ChainedSubcommandValue using its canonical wire layout.
func (x *ChainedSubcommandValue) Marshal(io IO) {
	io.Varuint32(&x.SubCommandFirstValue)
	io.Varuint32(&x.SubCommandSecondValue)
}

// Command holds the data that a command requires to be shown to a player client-side. The command
// is shown in the /help command and auto-completed using this data.
type Command struct {
	// Name is the name of the command. The command may be executed using this name, and will be shown
	// in the /help list with it. It currently seems that the client crashes if the Name contains
	// uppercase letters.
	Name string
	// Description is the description of the command. It is shown in the /help list and when starting to
	// write a command.
	Description string
	// Flags is a combination of flags not currently known. Leaving the Flags field empty appears to
	// work.
	Flags uint16
	// PermissionLevel is the command permission level that the player required to execute this command.
	// The field no longer seems to serve a purpose, as the client does not handle the execution of
	// commands anymore: The permissions should be checked server-side.
	PermissionLevel                     string
	AliasEnum                           int32
	CommandDataChainedSubcommandIndexes []uint32
	// Overloads is a list of command overloads that specify the ways in which a command may be
	// executed. The overloads may be completely different.
	Overloads []CommandOverload
}

// Marshal reads or writes Command using its canonical wire layout.
func (x *Command) Marshal(io IO) {
	io.String(&x.Name)
	io.String(&x.Description)
	io.Uint16(&x.Flags)
	io.String(&x.PermissionLevel)
	io.Int32(&x.AliasEnum)
	FuncSlice(io, &x.CommandDataChainedSubcommandIndexes, io.Varuint32, io.Uint32)
	Slice(io, &x.Overloads)
}

type CommandBlockUpdateData interface {
	isCommandBlockUpdateData()
}

// MarshalCommandBlockUpdateData reads or writes the CommandBlockUpdateData union using its canonical wire layout.
func MarshalCommandBlockUpdateData(io IO, x *CommandBlockUpdateData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(EntityCommandTarget)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(BlockCommandData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *EntityCommandTarget:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BlockCommandData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

// CommandEnum represents an enum in a command usage. The enum typically has a type and a set of
// options that are valid. A value that is not one of the options results in a failure during
// execution.
type CommandEnum struct {
	// Name is the type of the command enum. The type will show up in the command usage as the type of
	// the argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	Name string
	// Values holds a list of indices that point to the EnumValues slice in the AvailableCommandsPacket.
	// These represent the options of the enum.
	Values []uint32
}

// Marshal reads or writes CommandEnum using its canonical wire layout.
func (x *CommandEnum) Marshal(io IO) {
	io.String(&x.Name)
	FuncSlice(io, &x.Values, io.Varuint32, io.Uint32)
}

// CommandEnumConstraint is sent in the AvailableCommands packet to limit what values of an enum may
// be used taking in account things such as whether cheats are enabled.
type CommandEnumConstraint struct {
	// EnumValueSymbol points to an enum value in the AvailableCommands packet that this constraint
	// should apply to.
	EnumValueSymbol uint32
	// EnumSymbol points to an enum in the AvailableCommands packet to which this constraint should
	// apply to.
	EnumSymbol uint32
	// ConstraintIndices holds a slice of constraints as present in the constants above.
	ConstraintIndices []uint8
}

// Marshal reads or writes CommandEnumConstraint using its canonical wire layout.
func (x *CommandEnumConstraint) Marshal(io IO) {
	io.Uint32(&x.EnumValueSymbol)
	io.Uint32(&x.EnumSymbol)
	FuncSlice(io, &x.ConstraintIndices, io.Varuint32, io.Uint8)
}

// CommandOrigin holds data that identifies the origin of the requesting of a command. It holds
// several fields that may be used to get specific information. When sent in a CommandRequest
// packet, the same CommandOrigin should be sent in a CommandOutput packet.
type CommandOriginData struct {
	Type string
	// UUID is a unique identifier for every instantiation of a command.
	UUID uuid.UUID
	// RequestID is an ID that identifies the request of the client. The server should send a
	// CommandOrigin with the same request ID to ensure it can be matched with the request by the caller
	// of the command. This is especially important for websocket servers and it seems that this field
	// is only non-empty for these websocket servers.
	RequestID string
	PlayerID  int64
}

// Marshal reads or writes CommandOriginData using its canonical wire layout.
func (x *CommandOriginData) Marshal(io IO) {
	io.String(&x.Type)
	io.UUID(&x.UUID)
	io.String(&x.RequestID)
	io.Int64(&x.PlayerID)
}

type CommandOutputData struct {
	OutputType     string
	SuccessCount   uint32
	OutputMessages []CommandOutputMessage
	DataSet        Optional[string]
}

// Marshal reads or writes CommandOutputData using its canonical wire layout.
func (x *CommandOutputData) Marshal(io IO) {
	io.String(&x.OutputType)
	io.Uint32(&x.SuccessCount)
	Slice(io, &x.OutputMessages)
	OptionalFunc(io, &x.DataSet, io.String)
}

// CommandOutputMessage represents a message sent by a command that holds the output of one of the
// commands executed.
type CommandOutputMessage struct {
	MessageID  string
	Successful bool
	// Parameters is a list of parameters that serve to supply the message sent with additional
	// information, such as the position that a player was teleported to or the effect that was applied
	// to an entity. These parameters only apply for the Minecraft built-in command output.
	Parameters []string
}

// Marshal reads or writes CommandOutputMessage using its canonical wire layout.
func (x *CommandOutputMessage) Marshal(io IO) {
	io.String(&x.MessageID)
	io.Bool(&x.Successful)
	FuncSlice(io, &x.Parameters, io.Varuint32, io.String)
}

// CommandOverload represents an overload of a command. This overload can be compared to function
// overloading in languages such as java. It represents a single usage of the command. A command may
// have multiple different overloads, which are handled differently.
type CommandOverload struct {
	// IsChaining determines if the parameters use chained subcommands or not.
	IsChaining bool
	// ParameterData is a list of command parameters that are part of the overload. These parameters
	// specify the usage of the command when this overload is applied.
	ParameterData []CommandParameter
}

// Marshal reads or writes CommandOverload using its canonical wire layout.
func (x *CommandOverload) Marshal(io IO) {
	io.Bool(&x.IsChaining)
	Slice(io, &x.ParameterData)
}

// CommandParameter represents a single parameter of a command overload, which accepts a certain
// type of input values. It has a name and a type which show up client-side when a player is
// entering the command.
type CommandParameter struct {
	// Name is the name of the command parameter. It shows up in the usage like <$Name: $Type>, with the
	// exception of enum types, which show up simply as a list of options if the list is short enough
	// and Options is set to false.
	Name        string
	ParseSymbol uint32
	IsOptional  bool
	// Options holds a combinations of options that additionally apply to the command parameter. The
	// list of options can be found above.
	Options uint8
}

// Marshal reads or writes CommandParameter using its canonical wire layout.
func (x *CommandParameter) Marshal(io IO) {
	io.String(&x.Name)
	io.Uint32(&x.ParseSymbol)
	io.Bool(&x.IsOptional)
	io.Uint8(&x.Options)
}

type CommandPermissionLevel uint8

const (
	CommandPermissionLevelAny           CommandPermissionLevel = 0
	CommandPermissionLevelGameDirectors CommandPermissionLevel = 1
	CommandPermissionLevelAdmin         CommandPermissionLevel = 2
	CommandPermissionLevelHost          CommandPermissionLevel = 3
	CommandPermissionLevelOwner         CommandPermissionLevel = 4
	CommandPermissionLevelInternal      CommandPermissionLevel = 5
)

// DynamicEnum is an enum variant that can have its options changed during runtime, without sending
// a new AvailableCommands packet.
type DynamicEnum struct {
	// EnumName is the type of the command enum. The type will show up in the command usage as the type
	// of the argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	EnumName string
	// EnumOptions is a slice of possible options for the enum.
	EnumOptions []string
}

// Marshal reads or writes DynamicEnum using its canonical wire layout.
func (x *DynamicEnum) Marshal(io IO) {
	io.String(&x.EnumName)
	FuncSlice(io, &x.EnumOptions, io.Varuint32, io.String)
}
