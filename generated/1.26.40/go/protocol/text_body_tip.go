// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyTip struct {
	Value TextMessageOnly
}

func (*TextBodyTip) isTextBody() {}

// Marshal reads or writes TextBodyTip using its canonical wire layout.
func (x *TextBodyTip) Marshal(io IO) {
	x.Value.Marshal(io)
}
