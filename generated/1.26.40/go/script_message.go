// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ScriptMessage struct {
	MessageId    string
	MessageValue string
}

// Marshal reads or writes ScriptMessage using its canonical wire layout.
func (x *ScriptMessage) Marshal(io IO) {
	io.String(&x.MessageId)
	io.String(&x.MessageValue)
}
