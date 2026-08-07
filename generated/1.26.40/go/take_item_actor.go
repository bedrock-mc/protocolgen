// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TakeItemActor struct {
	ItemRuntimeID  uint64
	ActorRuntimeID uint64
}

// Marshal reads or writes TakeItemActor using its canonical wire layout.
func (x *TakeItemActor) Marshal(io IO) {
	io.ActorRuntimeID(&x.ItemRuntimeID)
	io.ActorRuntimeID(&x.ActorRuntimeID)
}
