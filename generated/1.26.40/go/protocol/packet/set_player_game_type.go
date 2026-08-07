// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetPlayerGameType struct {
	PlayerGameType protocol.GameType
}

// Marshal reads or writes SetPlayerGameType using its canonical wire layout.
func (x *SetPlayerGameType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PlayerGameType, io.Varint32)
}

// ID returns the protocol ID for SetPlayerGameType.
func (*SetPlayerGameType) ID() uint32 { return IDSetPlayerGameType }
