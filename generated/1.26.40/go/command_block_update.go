// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandBlockUpdate struct {
	Target             CommandBlockUpdateTarget
	Command            string
	LastOutput         string
	Name               string
	FilteredName       string
	TrackOutput        bool
	TickDelay          int32
	ExecuteOnFirstTick bool
}

// Marshal reads or writes CommandBlockUpdate using its canonical wire layout.
func (x *CommandBlockUpdate) Marshal(io IO) {
	marshalCommandBlockUpdateTarget(io, &x.Target)
	io.String(&x.Command)
	io.String(&x.LastOutput)
	io.String(&x.Name)
	io.String(&x.FilteredName)
	io.Bool(&x.TrackOutput)
	io.Int32(&x.TickDelay)
	io.Bool(&x.ExecuteOnFirstTick)
}
