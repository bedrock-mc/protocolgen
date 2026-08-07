// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SlashCommand struct {
	SuccessCount int32
	ErrorCount   int32
	CommandName  string
	ErrorList    string
}

func (*SlashCommand) isEventData() {}

// Marshal reads or writes SlashCommand using its canonical wire layout.
func (x *SlashCommand) Marshal(io IO) {
	io.Varint32(&x.SuccessCount)
	io.Varint32(&x.ErrorCount)
	io.String(&x.CommandName)
	io.String(&x.ErrorList)
}
