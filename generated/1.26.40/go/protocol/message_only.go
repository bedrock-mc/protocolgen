// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MessageOnly struct {
	Message string
}

func (*MessageOnly) isTextData() {}

// Marshal reads or writes MessageOnly using its canonical wire layout.
func (x *MessageOnly) Marshal(io IO) {
	io.String(&x.Message)
}
