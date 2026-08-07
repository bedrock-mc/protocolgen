// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetCommandsEnabled struct {
	CommandsEnabled bool
}

// Marshal reads or writes SetCommandsEnabled using its canonical wire layout.
func (x *SetCommandsEnabled) Marshal(io protocol.IO) {
	io.Bool(&x.CommandsEnabled)
}

// ID returns the protocol ID for SetCommandsEnabled.
func (*SetCommandsEnabled) ID() uint32 { return IDSetCommandsEnabled }
