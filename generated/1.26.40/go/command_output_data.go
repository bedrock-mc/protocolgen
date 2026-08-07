// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

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
	FuncSlice(io, &x.OutputMessages, io.Varuint32, func(value *CommandOutputMessage) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.DataSet, io.String)
}
