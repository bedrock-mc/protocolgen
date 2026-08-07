// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AgentAnimation struct {
	AgentAnimation protocol.AgentAnimationType
	RuntimeId      uint64
}

// Marshal reads or writes AgentAnimation using its canonical wire layout.
func (x *AgentAnimation) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.AgentAnimation, io.Uint8)
	io.ActorRuntimeID(&x.RuntimeId)
}

// ID returns the protocol ID for AgentAnimation.
func (*AgentAnimation) ID() uint32 { return IDAgentAnimation }
