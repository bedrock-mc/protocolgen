// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataSystemMessage struct {
	Value MessageOnly
}

func (*TextDataSystemMessage) isTextData() {}

// Marshal reads or writes TextDataSystemMessage using its canonical wire layout.
func (x *TextDataSystemMessage) Marshal(io IO) {
	x.Value.Marshal(io)
}
