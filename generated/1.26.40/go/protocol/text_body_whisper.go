// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyWhisper struct {
	Value TextAuthorAndMessage
}

func (*TextBodyWhisper) isTextBody() {}

// Marshal reads or writes TextBodyWhisper using its canonical wire layout.
func (x *TextBodyWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}
