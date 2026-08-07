// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SimpleEvent struct {
	Type protocol.SimpleEventSubtype
}

// Marshal reads or writes SimpleEvent using its canonical wire layout.
func (x *SimpleEvent) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Type, io.Uint16)
}
