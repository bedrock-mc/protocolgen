// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SimpleEvent struct {
	Type protocol.Subtype
}

// Marshal reads or writes SimpleEvent using its canonical wire layout.
func (x *SimpleEvent) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Type, io.Uint16)
}

// ID returns the protocol ID for SimpleEvent.
func (*SimpleEvent) ID() uint32 { return IDSimpleEvent }
