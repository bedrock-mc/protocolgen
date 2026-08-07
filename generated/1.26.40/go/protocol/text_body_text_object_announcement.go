// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyTextObjectAnnouncement struct {
	Value TextMessageOnly
}

func (TextBodyTextObjectAnnouncement) isTextBody() {}

// Marshal reads or writes TextBodyTextObjectAnnouncement using its canonical wire layout.
func (x *TextBodyTextObjectAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}
