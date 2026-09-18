// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
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

// ID ...
func (*ClientboundTextureShift) ID() uint32 {
	return IDClientboundTextureShift
}

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
