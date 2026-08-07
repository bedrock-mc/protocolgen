// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerFog struct {
	FogStack []string
}

// Marshal reads or writes PlayerFog using its canonical wire layout.
func (x *PlayerFog) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.FogStack, io.Varuint32, io.String)
}

// ID returns the protocol ID for PlayerFog.
func (*PlayerFog) ID() uint32 { return IDPlayerFog }
