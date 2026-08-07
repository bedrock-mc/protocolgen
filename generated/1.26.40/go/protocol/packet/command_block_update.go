// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CommandBlockUpdate struct {
	Target             protocol.CommandBlockUpdateTarget
	Command            string
	LastOutput         string
	Name               string
	FilteredName       string
	TrackOutput        bool
	TickDelay          int32
	ExecuteOnFirstTick bool
}

// Marshal reads or writes CommandBlockUpdate using its canonical wire layout.
func (x *CommandBlockUpdate) Marshal(io protocol.IO) {
	protocol.MarshalCommandBlockUpdateTarget(io, &x.Target)
	io.String(&x.Command)
	io.String(&x.LastOutput)
	io.String(&x.Name)
	io.String(&x.FilteredName)
	io.Bool(&x.TrackOutput)
	io.Int32(&x.TickDelay)
	io.Bool(&x.ExecuteOnFirstTick)
}

// ID returns the protocol ID for CommandBlockUpdate.
func (*CommandBlockUpdate) ID() uint32 { return IDCommandBlockUpdate }
