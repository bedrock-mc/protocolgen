// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type Command struct {
	Name                                string
	Description                         string
	Flags                               uint16
	PermissionLevel                     string
	AliasEnum                           int32
	CommandDataChainedSubcommandIndexes []uint32
	Overloads                           []CommandOverload
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
