// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AgentActionEvent struct {
	RequestId string
	Action    AgentActionType
	Response  string
}

// Marshal reads or writes AgentActionEvent using its canonical wire layout.
func (x *AgentActionEvent) Marshal(io IO) {
	io.String(&x.RequestId)
	IntegerFunc(&x.Action, io.Int32)
	io.String(&x.Response)
}
