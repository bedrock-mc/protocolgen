// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MessageAndParams struct {
	Message       string
	ParameterList []string
}

func (*MessageAndParams) isTextData() {}

// Marshal reads or writes MessageAndParams using its canonical wire layout.
func (x *MessageAndParams) Marshal(io IO) {
	io.String(&x.Message)
	FuncSlice(io, &x.ParameterList, io.Varuint32, io.String)
}
