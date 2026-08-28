// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// TakeItemActor is sent by the server when a player picks up an item entity. It makes the item
// entity disappear to viewers and shows the pick-up animation.
type TakeItemActor struct {
	ItemRuntimeID  uint64
	ActorRuntimeID uint64
}

// Marshal reads or writes TakeItemActor using its canonical wire layout.
func (x *TakeItemActor) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.ItemRuntimeID)
	io.ActorRuntimeID(&x.ActorRuntimeID)
}

// ID returns the protocol ID for TakeItemActor.
func (*TakeItemActor) ID() uint32 { return IDTakeItemActor }
