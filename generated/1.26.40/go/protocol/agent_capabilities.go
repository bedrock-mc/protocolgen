// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AgentCapabilities struct {
	CanModifyBlocks Optional[bool]
}

// Marshal reads or writes AgentCapabilities using its canonical wire layout.
func (x *AgentCapabilities) Marshal(io IO) {
	OptionalFunc(io, &x.CanModifyBlocks, io.Bool)
}
