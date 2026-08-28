// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// AutomationClientConnect is used to make the client connect to a websocket server. This websocket
// server has the ability to execute commands on the behalf of the client and it can listen for
// certain events fired by the client.
type AutomationClientConnect struct {
	WebSocketData protocol.WebSocketData
}

// Marshal reads or writes AutomationClientConnect using its canonical wire layout.
func (x *AutomationClientConnect) Marshal(io protocol.IO) {
	x.WebSocketData.Marshal(io)
}

// ID returns the protocol ID for AutomationClientConnect.
func (*AutomationClientConnect) ID() uint32 { return IDAutomationClientConnect }
