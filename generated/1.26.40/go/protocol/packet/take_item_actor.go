// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type TakeItemActor struct {
	ItemRuntimeID  uint64
	ActorRuntimeID uint64
}

// Marshal reads or writes TakeItemActor using its canonical wire layout.
func (x *TakeItemActor) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.ItemRuntimeID)
	io.ActorRuntimeID(&x.ActorRuntimeID)
}
