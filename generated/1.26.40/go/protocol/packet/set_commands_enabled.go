// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SetCommandsEnabled is sent by the server to enable or disable the ability to execute commands for
// the client. If disabled, the client itself will stop the execution of commands.
type SetCommandsEnabled struct {
	// Enabled defines if the commands should be enabled, or if false, disabled.
	CommandsEnabled bool
}

// Marshal reads or writes SetCommandsEnabled using its canonical wire layout.
func (x *SetCommandsEnabled) Marshal(io protocol.IO) {
	io.Bool(&x.CommandsEnabled)
}

// ID returns the protocol ID for SetCommandsEnabled.
func (*SetCommandsEnabled) ID() uint32 { return IDSetCommandsEnabled }
