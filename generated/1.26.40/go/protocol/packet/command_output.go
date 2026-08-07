// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CommandOutput struct {
	OriginData protocol.CommandOriginData
	Output     protocol.CommandOutputData
}

// Marshal reads or writes CommandOutput using its canonical wire layout.
func (x *CommandOutput) Marshal(io protocol.IO) {
	x.OriginData.Marshal(io)
	x.Output.Marshal(io)
}
