// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ActorPickRequest struct {
	ActorID  int64
	MaxSlots uint8
	WithData bool
}

// Marshal reads or writes ActorPickRequest using its canonical wire layout.
func (x *ActorPickRequest) Marshal(io IO) {
	io.Int64(&x.ActorID)
	io.Uint8(&x.MaxSlots)
	io.Bool(&x.WithData)
}
