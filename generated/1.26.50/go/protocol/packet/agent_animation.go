// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// AgentAnimation is an Education Edition packet sent from the server to the client to make an agent
// perform an animation.
type AgentAnimation struct {
	AgentAnimation protocol.AgentAnimationType
	RuntimeID      uint64
}

// Marshal reads or writes AgentAnimation using its canonical wire layout.
func (x *AgentAnimation) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.AgentAnimation, io.Uint8)
	io.ActorRuntimeID(&x.RuntimeID)
}

// ID returns the protocol ID for AgentAnimation.
func (*AgentAnimation) ID() uint32 { return IDAgentAnimation }
