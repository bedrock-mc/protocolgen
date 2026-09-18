// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// TakeItemActor is sent by the server when a player picks up an item entity. It makes the item entity
// disappear to viewers and shows the pick-up animation.
type TakeItemActor struct {
	ItemRuntimeID  uint64
	ActorRuntimeID uint64
}

// ID ...
func (*TakeItemActor) ID() uint32 {
	return IDTakeItemActor
}

func (pk *TakeItemActor) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.ItemRuntimeID)
	io.ActorRuntimeID(&pk.ActorRuntimeID)
}
