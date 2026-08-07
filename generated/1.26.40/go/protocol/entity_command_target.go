// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EntityCommandTarget struct {
	TargetRuntimeID uint64
}

func (*EntityCommandTarget) isCommandBlockUpdateData() {}

// Marshal reads or writes EntityCommandTarget using its canonical wire layout.
func (x *EntityCommandTarget) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
}
