// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CommandRequest is sent by the client to request the execution of a server-side command. Although
// some servers support sending commands using the Text packet, this packet is guaranteed to have
// the correct result.
type CommandRequest struct {
	// Command is the raw entered command line. The client does no parsing of the command line by itself
	// (unlike it did in the early stages), but lets the server do that.
	Command string
	// Origin is the data specifying the origin of the command. In other words, the source that the
	// command was from, such as the player itself or a websocket server.
	Origin protocol.CommandOriginData
	// IsInternal specifies if the command request internal. Setting it to false seems to work and the
	// usage of this field is not known.
	IsInternal bool
	// Version is the version of the command that is being executed. This field currently has no purpose
	// or functionality.
	Version string
}

// Marshal reads or writes CommandRequest using its canonical wire layout.
func (x *CommandRequest) Marshal(io protocol.IO) {
	io.StringLimits(&x.Command, 0, 1000)
	x.Origin.Marshal(io)
	io.Bool(&x.IsInternal)
	io.String(&x.Version)
}

// ID returns the protocol ID for CommandRequest.
func (*CommandRequest) ID() uint32 { return IDCommandRequest }
