// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Disconnect struct {
	Reason   ConnectionDisconnectFailReason
	Messages DisconnectMessages
}

// Marshal reads or writes Disconnect using its canonical wire layout.
func (x *Disconnect) Marshal(io IO) {
	IntegerFunc(&x.Reason, io.Varint32)
	marshalDisconnectMessages(io, &x.Messages)
}
