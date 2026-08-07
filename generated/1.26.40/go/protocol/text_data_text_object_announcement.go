// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataTextObjectAnnouncement struct {
	Value MessageOnly
}

func (*TextDataTextObjectAnnouncement) isTextData() {}

// Marshal reads or writes TextDataTextObjectAnnouncement using its canonical wire layout.
func (x *TextDataTextObjectAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}
