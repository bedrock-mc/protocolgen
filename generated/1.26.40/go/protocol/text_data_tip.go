// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataTip struct {
	Value MessageOnly
}

func (*TextDataTip) isTextData() {}

// Marshal reads or writes TextDataTip using its canonical wire layout.
func (x *TextDataTip) Marshal(io IO) {
	x.Value.Marshal(io)
}
