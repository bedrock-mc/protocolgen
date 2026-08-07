// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyPopup struct {
	Value TextMessageAndParams
}

func (*TextBodyPopup) isTextBody() {}

// Marshal reads or writes TextBodyPopup using its canonical wire layout.
func (x *TextBodyPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}
