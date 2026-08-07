// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AgentActionEvent struct {
	RequestID string
	Action    protocol.AgentActionType
	Response  string
}

// Marshal reads or writes AgentActionEvent using its canonical wire layout.
func (x *AgentActionEvent) Marshal(io protocol.IO) {
	io.String(&x.RequestID)
	protocol.IntegerFunc(&x.Action, io.Int32)
	io.String(&x.Response)
}

// ID returns the protocol ID for AgentActionEvent.
func (*AgentActionEvent) ID() uint32 { return IDAgentActionEvent }
