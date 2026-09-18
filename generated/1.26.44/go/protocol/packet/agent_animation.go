// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// AgentAnimation is an Education Edition packet sent from the server to the client to make an agent perform
// an animation.
type AgentAnimation struct {
	AgentAnimation protocol.AgentAnimationType
	RuntimeID      uint64
}

// ID ...
func (*AgentAnimation) ID() uint32 {
	return IDAgentAnimation
}

func (pk *AgentAnimation) Marshal(io protocol.IO) {
	pk.AgentAnimation.Marshal(io)
	io.ActorRuntimeID(&pk.RuntimeID)
}
