// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AuthorAndMessage struct {
	PlayerName string
	Message    string
}

func (*AuthorAndMessage) isTextData() {}

// Marshal reads or writes AuthorAndMessage using its canonical wire layout.
func (x *AuthorAndMessage) Marshal(io IO) {
	io.String(&x.PlayerName)
	io.String(&x.Message)
}
