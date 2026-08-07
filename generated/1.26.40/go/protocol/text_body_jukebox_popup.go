// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyJukeboxPopup struct {
	Value TextMessageAndParams
}

func (TextBodyJukeboxPopup) isTextBody() {}

// Marshal reads or writes TextBodyJukeboxPopup using its canonical wire layout.
func (x *TextBodyJukeboxPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}
