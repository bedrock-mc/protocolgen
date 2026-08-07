// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TextMessageAndParams struct {
	Message       string
	ParameterList []string
}

func (TextMessageAndParams) isTextBody() {}

// Marshal reads or writes TextMessageAndParams using its canonical wire layout.
func (x *TextMessageAndParams) Marshal(io IO) {
	io.String(&x.Message)
	FuncSlice(io, &x.ParameterList, io.Varuint32, io.String)
}
