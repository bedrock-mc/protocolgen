// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandOutputMessage struct {
	MessageID  string
	Successful bool
	Parameters []string
}

// Marshal reads or writes CommandOutputMessage using its canonical wire layout.
func (x *CommandOutputMessage) Marshal(io IO) {
	io.String(&x.MessageID)
	io.Bool(&x.Successful)
	FuncSlice(io, &x.Parameters, io.Varuint32, io.String)
}
