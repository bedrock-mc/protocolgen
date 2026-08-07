// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataPopup struct {
	Value MessageAndParams
}

func (*TextDataPopup) isTextData() {}

// Marshal reads or writes TextDataPopup using its canonical wire layout.
func (x *TextDataPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}
