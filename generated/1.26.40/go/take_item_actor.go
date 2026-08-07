// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TakeItemActor struct {
	ItemRuntimeID  ActorRuntimeID
	ActorRuntimeID ActorRuntimeID
}

// Marshal reads or writes TakeItemActor using its canonical wire layout.
func (x *TakeItemActor) Marshal(io IO) {
	x.ItemRuntimeID.Marshal(io)
	x.ActorRuntimeID.Marshal(io)
}
