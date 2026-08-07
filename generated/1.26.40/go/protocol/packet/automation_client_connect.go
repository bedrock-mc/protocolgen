// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AutomationClientConnect struct {
	WebSocketData protocol.WebSocketPacketData
}

// Marshal reads or writes AutomationClientConnect using its canonical wire layout.
func (x *AutomationClientConnect) Marshal(io protocol.IO) {
	x.WebSocketData.Marshal(io)
}
