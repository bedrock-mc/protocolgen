// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TextBodyTextObject struct {
	Value TextMessageOnly
}

func (TextBodyTextObject) isTextBody() {}

// Marshal reads or writes TextBodyTextObject using its canonical wire layout.
func (x *TextBodyTextObject) Marshal(io IO) {
	x.Value.Marshal(io)
}
