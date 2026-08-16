// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ScriptMessage is used to communicate custom messages from the client to the server, or from the
// server to the client. While the name may suggest this packet is used for the discontinued
// scripting API, it is likely instead for the GameTest framework.
type ScriptMessage struct {
	MessageID    string
	MessageValue []byte
}

// Marshal reads or writes ScriptMessage using its canonical wire layout.
func (x *ScriptMessage) Marshal(io protocol.IO) {
	io.String(&x.MessageID)
	io.Bytes(&x.MessageValue)
}

// ID returns the protocol ID for ScriptMessage.
func (*ScriptMessage) ID() uint32 { return IDScriptMessage }
