// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBodyAnnouncement struct {
	Value TextAuthorAndMessage
}

func (*TextBodyAnnouncement) isTextBody() {}

// Marshal reads or writes TextBodyAnnouncement using its canonical wire layout.
func (x *TextBodyAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}
