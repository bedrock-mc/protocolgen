// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// RemoveActor is sent by the server to remove an entity that currently exists in the world from the client-
// side. Sending this packet if the client cannot already see this entity will have no effect.
type RemoveActor struct {
	TargetActorID int64
}

// ID ...
func (*RemoveActor) ID() uint32 {
	return IDRemoveActor
}

func (pk *RemoveActor) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.TargetActorID)
}
