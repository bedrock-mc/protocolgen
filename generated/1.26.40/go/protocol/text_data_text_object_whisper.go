// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataTextObjectWhisper struct {
	Value MessageOnly
}

func (*TextDataTextObjectWhisper) isTextData() {}

// Marshal reads or writes TextDataTextObjectWhisper using its canonical wire layout.
func (x *TextDataTextObjectWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}
