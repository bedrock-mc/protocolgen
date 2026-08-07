// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Animate struct {
	Action               AnimateAction
	TargetActorRuntimeID uint64
	Data                 float32
	SwingSource          Optional[string]
}

// Marshal reads or writes Animate using its canonical wire layout.
func (x *Animate) Marshal(io IO) {
	IntegerFunc(&x.Action, io.Uint8)
	io.ActorRuntimeID(&x.TargetActorRuntimeID)
	io.Float32(&x.Data)
	OptionalFunc(io, &x.SwingSource, io.String)
}
