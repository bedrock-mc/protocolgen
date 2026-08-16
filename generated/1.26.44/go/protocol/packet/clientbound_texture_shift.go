// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

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

// Marshal reads or writes ClientboundTextureShift using its canonical wire layout.
func (x *ClientboundTextureShift) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.ActionID, io.Uint8)
	io.String(&x.CollectionName)
	io.String(&x.FromStep)
	io.String(&x.ToStep)
	protocol.FuncSlice(io, &x.AllSteps, io.Varuint32, io.String)
	io.Varuint64(&x.CurrentLengthInTicks)
	protocol.Minimum(io, &x.CurrentLengthInTicks, 0)
	io.Varuint64(&x.TotalLengthInTicks)
	protocol.Minimum(io, &x.TotalLengthInTicks, 0)
	io.Bool(&x.Enabled)
}

// ID returns the protocol ID for ClientboundTextureShift.
func (*ClientboundTextureShift) ID() uint32 { return IDClientboundTextureShift }
