// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// DebugInfo is a packet sent by the server to the client. It does not seem to do anything when sent to the
// normal client in 1.16.
type DebugInfo struct {
	ActorID int64
	// Data is the debug data.
	Data []byte
}

// ID ...
func (*DebugInfo) ID() uint32 {
	return IDDebugInfo
}

func (pk *DebugInfo) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.ActorID)
	io.ByteSlice(&pk.Data)
}
