// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextDataAnnouncement struct {
	Value AuthorAndMessage
}

func (*TextDataAnnouncement) isTextData() {}

// Marshal reads or writes TextDataAnnouncement using its canonical wire layout.
func (x *TextDataAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}
