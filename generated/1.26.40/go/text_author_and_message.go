// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TextAuthorAndMessage struct {
	PlayerName string
	Message    string
}

func (TextAuthorAndMessage) isTextBody() {}

// Marshal reads or writes TextAuthorAndMessage using its canonical wire layout.
func (x *TextAuthorAndMessage) Marshal(io IO) {
	io.String(&x.PlayerName)
	io.String(&x.Message)
}
