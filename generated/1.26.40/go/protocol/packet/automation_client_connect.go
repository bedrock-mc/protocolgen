// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// AutomationClientConnect is used to make the client connect to a websocket server. This websocket server has
// the ability to execute commands on the behalf of the client and it can listen for certain events fired by
// the client.
type AutomationClientConnect struct {
	WebsocketServerURI string
}

// ID ...
func (*AutomationClientConnect) ID() uint32 {
	return IDAutomationClientConnect
}

func (pk *AutomationClientConnect) Marshal(io protocol.IO) {
	io.String(&pk.WebsocketServerURI)
}
