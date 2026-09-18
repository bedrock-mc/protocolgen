// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type AgentActionEvent struct {
	RequestID string
	Action    protocol.AgentActionType
	Response  string
}

// ID ...
func (*AgentActionEvent) ID() uint32 {
	return IDAgentActionEvent
}

func (pk *AgentActionEvent) Marshal(io protocol.IO) {
	io.String(&pk.RequestID)
	pk.Action.Marshal(io)
	io.String(&pk.Response)
}
