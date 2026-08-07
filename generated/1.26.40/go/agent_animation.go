// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AgentAnimation struct {
	AgentAnimation AgentAnimationType
	RuntimeId      uint64
}

// Marshal reads or writes AgentAnimation using its canonical wire layout.
func (x *AgentAnimation) Marshal(io IO) {
	IntegerFunc(&x.AgentAnimation, io.Uint8)
	io.ActorRuntimeID(&x.RuntimeId)
}
