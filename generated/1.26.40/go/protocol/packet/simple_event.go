// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SimpleEvent is used for enabling or disabling commands and for unlocking world template settings
// (both unlocking UI buttons on client and the actual setting on the server). This is fired from
// the client to the server and a SetCommandsEnabled is sent back when enabling commands.
type SimpleEvent struct {
	Type protocol.Subtype
}

// Marshal reads or writes SimpleEvent using its canonical wire layout.
func (x *SimpleEvent) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Type, io.Uint16)
}

// ID returns the protocol ID for SimpleEvent.
func (*SimpleEvent) ID() uint32 { return IDSimpleEvent }
