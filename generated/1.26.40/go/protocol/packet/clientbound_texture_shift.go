// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ClientboundTextureShift struct {
	ActionID             protocol.ClientboundTextureShiftAction
	CollectionName       string
	FromStep             string
	ToStep               string
	AllSteps             []string
	CurrentLengthInTicks uint64
	TotalLengthInTicks   uint64
	Enabled              bool
}

// ID returns the protocol ID for ClientboundTextureShift.
func (*ClientboundTextureShift) ID() uint32 { return IDClientboundTextureShift }

// Marshal reads or writes ClientboundTextureShift using its canonical wire layout.
func (pk *ClientboundTextureShift) Marshal(io protocol.IO) {
	pk.ActionID.Marshal(io)
	io.String(&pk.CollectionName)
	io.String(&pk.FromStep)
	io.String(&pk.ToStep)
	protocol.FuncSlice(io, &pk.AllSteps, io.Varuint32, io.String)
	io.Varuint64(&pk.CurrentLengthInTicks)
	io.Varuint64(&pk.TotalLengthInTicks)
	io.Bool(&pk.Enabled)
}
