// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyTextObjectWhisper struct {
	Value TextMessageOnly
}

func (*TextBodyTextObjectWhisper) isTextBody() {}

// Marshal reads or writes TextBodyTextObjectWhisper using its canonical wire layout.
func (x *TextBodyTextObjectWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}
