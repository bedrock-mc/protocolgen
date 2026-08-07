// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataWhisper struct {
	Value AuthorAndMessage
}

func (*TextDataWhisper) isTextData() {}

// Marshal reads or writes TextDataWhisper using its canonical wire layout.
func (x *TextDataWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}
