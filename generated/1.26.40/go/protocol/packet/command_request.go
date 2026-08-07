// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CommandRequest struct {
	Command    string
	Origin     protocol.CommandOriginData
	IsInternal bool
	Version    string
}

// Marshal reads or writes CommandRequest using its canonical wire layout.
func (x *CommandRequest) Marshal(io protocol.IO) {
	io.String(&x.Command)
	x.Origin.Marshal(io)
	io.Bool(&x.IsInternal)
	io.String(&x.Version)
}

// ID returns the protocol ID for CommandRequest.
func (*CommandRequest) ID() uint32 { return IDCommandRequest }
