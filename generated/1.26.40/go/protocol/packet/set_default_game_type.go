// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetDefaultGameType struct {
	DefaultGameType protocol.GameType
}

// Marshal reads or writes SetDefaultGameType using its canonical wire layout.
func (x *SetDefaultGameType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.DefaultGameType, io.Varint32)
}

// ID returns the protocol ID for SetDefaultGameType.
func (*SetDefaultGameType) ID() uint32 { return IDSetDefaultGameType }
