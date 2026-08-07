// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ScriptMessage struct {
	MessageId    string
	MessageValue string
}

// Marshal reads or writes ScriptMessage using its canonical wire layout.
func (x *ScriptMessage) Marshal(io protocol.IO) {
	io.String(&x.MessageId)
	io.String(&x.MessageValue)
}

// ID returns the protocol ID for ScriptMessage.
func (*ScriptMessage) ID() uint32 { return IDScriptMessage }
