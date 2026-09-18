// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// AgentActionEvent is an Education Edition packet sent from the server to the client to return a response to
// a previously requested action.
type AgentActionEvent struct {
	// Identifier is a JSON identifier referenced in the initial action.
	RequestID string
	// Action represents the action type that was requested. It is one of the constants defined above.
	Action protocol.AgentActionType
	// Response is a JSON string containing the response to the action.
	Response string
}

// ID returns the protocol ID for AgentActionEvent.
func (*AgentActionEvent) ID() uint32 { return IDAgentActionEvent }

// Marshal reads or writes AgentActionEvent using its canonical wire layout.
func (pk *AgentActionEvent) Marshal(io protocol.IO) {
	io.String(&pk.RequestID)
	pk.Action.Marshal(io)
	io.String(&pk.Response)
}
