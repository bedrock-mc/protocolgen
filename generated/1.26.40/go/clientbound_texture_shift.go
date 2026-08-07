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
	enumValue1 := uint8(x.ActionID)
	io.Uint8(&enumValue1)
	x.ActionID = ClientboundTextureShiftAction(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.CollectionName)
	io.String(&x.FromStep)
	io.String(&x.ToStep)
	if !io.Reading() && uint64(len(x.AllSteps)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.AllSteps), "collection length overflows uint32")
		return
	}
	count2 := uint32(len(x.AllSteps))
	io.Varuint32(&count2)
	if io.Reading() {
		if uint64(count2) > uint64(^uint(0)>>1) {
			io.InvalidValue(count2, "collection length overflows int")
			return
		}
		x.AllSteps = make([]string, int(count2))
	}
	for index3 := range x.AllSteps {
		io.String(&x.AllSteps[index3])
	}
	io.Varuint64(&x.CurrentLengthInTicks)
	io.Varuint64(&x.TotalLengthInTicks)
	io.Bool(&x.Enabled)
}
