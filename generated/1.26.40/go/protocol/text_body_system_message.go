// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodySystemMessage struct {
	Value TextMessageOnly
}

func (*TextBodySystemMessage) isTextBody() {}

// Marshal reads or writes TextBodySystemMessage using its canonical wire layout.
func (x *TextBodySystemMessage) Marshal(io IO) {
	x.Value.Marshal(io)
}
