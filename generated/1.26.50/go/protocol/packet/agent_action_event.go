// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type AgentActionEvent struct {
	RequestID string
	Action    protocol.AgentActionType
	Response  string
}

// ID returns the protocol ID for AgentActionEvent.
func (*AgentActionEvent) ID() uint32 { return IDAgentActionEvent }

// Marshal reads or writes AgentActionEvent using its canonical wire layout.
func (pk *AgentActionEvent) Marshal(io protocol.IO) {
	io.String(&pk.RequestID)
	pk.Action.Marshal(io)
	io.String(&pk.Response)
}
