// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type DebugInfo struct {
	ActorID int64
	Data    string
}

// Marshal reads or writes DebugInfo using its canonical wire layout.
func (x *DebugInfo) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.ActorID)
	io.String(&x.Data)
}

// ID returns the protocol ID for DebugInfo.
func (*DebugInfo) ID() uint32 { return IDDebugInfo }
