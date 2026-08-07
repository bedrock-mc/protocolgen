// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SetDefaultGameType is sent by the client when it toggles the default game type in the settings
// UI, and is sent by the server when it actually changes the default game type, resulting in the
// toggle being changed in the settings UI.
type SetDefaultGameType struct {
	DefaultGameType protocol.GameType
}

// Marshal reads or writes SetDefaultGameType using its canonical wire layout.
func (x *SetDefaultGameType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.DefaultGameType, io.Varint32)
}

// ID returns the protocol ID for SetDefaultGameType.
func (*SetDefaultGameType) ID() uint32 { return IDSetDefaultGameType }
