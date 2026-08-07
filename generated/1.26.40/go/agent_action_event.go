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
	enumValue1 := int32(x.Action)
	io.Int32(&enumValue1)
	x.Action = AgentActionType(enumValue1)
	switch int64(enumValue1) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.Response)
}
