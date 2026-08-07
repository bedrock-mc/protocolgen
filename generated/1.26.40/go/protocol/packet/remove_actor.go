// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RemoveActor struct {
	TargetActorID int64
}

// Marshal reads or writes RemoveActor using its canonical wire layout.
func (x *RemoveActor) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
}
