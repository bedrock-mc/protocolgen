// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DisconnectPacketMessages struct {
	Message         string
	FilteredMessage string
}

func (*DisconnectPacketMessages) isDisconnectMessages() {}

// Marshal reads or writes DisconnectPacketMessages using its canonical wire layout.
func (x *DisconnectPacketMessages) Marshal(io IO) {
	io.String(&x.Message)
	io.String(&x.FilteredMessage)
}
