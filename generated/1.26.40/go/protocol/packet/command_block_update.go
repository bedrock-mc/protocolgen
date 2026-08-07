// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CommandBlockUpdate is sent by the client to update a command block at a specific position. The
// command block may be either a physical block or an entity.
type CommandBlockUpdate struct {
	Target protocol.CommandBlockUpdateData
	// Command is the command currently entered in the command block. This is the command that is
	// executed when the command block is activated.
	Command string
	// LastOutput is the output of the last command executed by the command block. It may be left empty
	// to show simply no output at all, in combination with setting ShouldTrackOutput to false.
	LastOutput string
	// Name is the name of the command block updated. If not empty, it will show this name hovering
	// above the command block when hovering over the block with the cursor.
	Name string
	// FilteredName is a filtered version of Name with all the profanity removed. The client will use
	// this over Name if this field is not empty and they have the "Filter Profanity" setting enabled.
	FilteredName string
	TrackOutput  bool
	// TickDelay is the delay in ticks between executions of a command block, if it is a repeating
	// command block.
	TickDelay int32
	// ExecuteOnFirstTick specifies if the command block should execute on the first tick, AKA as soon
	// as the command block is enabled.
	ExecuteOnFirstTick bool
}

// Marshal reads or writes CommandBlockUpdate using its canonical wire layout.
func (x *CommandBlockUpdate) Marshal(io protocol.IO) {
	protocol.MarshalCommandBlockUpdateData(io, &x.Target)
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
