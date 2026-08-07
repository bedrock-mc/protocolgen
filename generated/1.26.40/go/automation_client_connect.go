// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AutomationClientConnect struct {
	WebSocketData WebSocketPacketData
}

// Marshal reads or writes AutomationClientConnect using its canonical wire layout.
func (x *AutomationClientConnect) Marshal(io IO) {
	x.WebSocketData.Marshal(io)
}
