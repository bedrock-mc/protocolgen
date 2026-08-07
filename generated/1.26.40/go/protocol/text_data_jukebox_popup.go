// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataJukeboxPopup struct {
	Value MessageAndParams
}

func (*TextDataJukeboxPopup) isTextData() {}

// Marshal reads or writes TextDataJukeboxPopup using its canonical wire layout.
func (x *TextDataJukeboxPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}
