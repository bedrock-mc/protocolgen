// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextMessageOnly struct {
	Message string
}

func (*TextMessageOnly) isTextBody() {}

// Marshal reads or writes TextMessageOnly using its canonical wire layout.
func (x *TextMessageOnly) Marshal(io IO) {
	io.String(&x.Message)
}
