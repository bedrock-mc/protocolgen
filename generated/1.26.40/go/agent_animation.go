// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AgentAnimation struct {
	AgentAnimation AgentAnimationType
	RuntimeId      ActorRuntimeID
}

// Marshal reads or writes AgentAnimation using its canonical wire layout.
func (x *AgentAnimation) Marshal(io IO) {
	enumValue1 := uint8(x.AgentAnimation)
	io.Uint8(&enumValue1)
	x.AgentAnimation = AgentAnimationType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.RuntimeId.Marshal(io)
}
