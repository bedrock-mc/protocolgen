// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundTextureShift struct {
	ActionID             ClientboundTextureShiftAction
	CollectionName       string
	FromStep             string
	ToStep               string
	AllSteps             []string
	CurrentLengthInTicks uint64
	TotalLengthInTicks   uint64
	Enabled              bool
}

// Marshal reads or writes ClientboundTextureShift using its canonical wire layout.
func (x *ClientboundTextureShift) Marshal(io IO) {
	IntegerFunc(&x.ActionID, io.Uint8)
	io.String(&x.CollectionName)
	io.String(&x.FromStep)
	io.String(&x.ToStep)
	FuncSlice(io, &x.AllSteps, io.Varuint32, io.String)
	io.Varuint64(&x.CurrentLengthInTicks)
	io.Varuint64(&x.TotalLengthInTicks)
	io.Bool(&x.Enabled)
}
