// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Animate struct {
	Action               AnimateAction
	TargetActorRuntimeID ActorRuntimeID
	Data                 float32
	SwingSource          Optional[string]
}

// Marshal reads or writes Animate using its canonical wire layout.
func (x *Animate) Marshal(io IO) {
	enumValue1 := uint8(x.Action)
	io.Uint8(&enumValue1)
	x.Action = AnimateAction(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 3, 4, 5:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.TargetActorRuntimeID.Marshal(io)
	io.Float32(&x.Data)
	io.Bool(&x.SwingSource.set)
	if x.SwingSource.set {
		io.String(&x.SwingSource.val)
	} else if io.Reading() {
		var zero string
		x.SwingSource.val = zero
	}
}
