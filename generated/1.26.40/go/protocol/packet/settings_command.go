// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SettingsCommand struct {
	Command        string
	SuppressOutput bool
}

// Marshal reads or writes SettingsCommand using its canonical wire layout.
func (x *SettingsCommand) Marshal(io protocol.IO) {
	io.String(&x.Command)
	io.Bool(&x.SuppressOutput)
}
