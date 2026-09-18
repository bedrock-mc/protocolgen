// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// AutomationClientConnect is used to make the client connect to a websocket server. This websocket server has
// the ability to execute commands on the behalf of the client and it can listen for certain events fired by
// the client.
type AutomationClientConnect struct {
	// ServerURI is the URI to make the client connect to. It can be, for example, 'localhost:8000/ws' to connect
	// to a websocket server on the localhost at port 8000.
	WebsocketServerURI string
}

// ID returns the protocol ID for AutomationClientConnect.
func (*AutomationClientConnect) ID() uint32 { return IDAutomationClientConnect }

// Marshal reads or writes AutomationClientConnect using its canonical wire layout.
func (pk *AutomationClientConnect) Marshal(io protocol.IO) {
	io.String(&pk.WebsocketServerURI)
}
