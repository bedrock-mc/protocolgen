// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// Animate is sent by the server to send a player animation from one player to all viewers of that player. It
// is used for a couple of actions, such as arm swimming and critical hits.
type Animate struct {
	Action               protocol.AnimateAction
	TargetActorRuntimeID uint64
	// Data ...
	Data float32
	// SwingSource is the source for swing actions. It is one of the action type constants that may be found
	// above.
	SwingSource protocol.Optional[string]
}

// ID ...
func (*Animate) ID() uint32 {
	return IDAnimate
}

func (pk *Animate) Marshal(io protocol.IO) {
	pk.Action.Marshal(io)
	io.ActorRuntimeID(&pk.TargetActorRuntimeID)
	io.Float32(&pk.Data)
	protocol.OptionalFunc(io, &pk.SwingSource, io.String)
}
