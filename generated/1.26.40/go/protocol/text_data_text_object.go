// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataTextObject struct {
	Value MessageOnly
}

func (*TextDataTextObject) isTextData() {}

// Marshal reads or writes TextDataTextObject using its canonical wire layout.
func (x *TextDataTextObject) Marshal(io IO) {
	x.Value.Marshal(io)
}
