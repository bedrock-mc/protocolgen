// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DisconnectMessagesData struct {
	Message         string
	FilteredMessage string
}

func (*DisconnectMessagesData) isDisconnectMessages() {}

// Marshal reads or writes DisconnectMessagesData using its canonical wire layout.
func (x *DisconnectMessagesData) Marshal(io IO) {
	io.String(&x.Message)
	io.String(&x.FilteredMessage)
}
