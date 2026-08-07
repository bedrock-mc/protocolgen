// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandRequest struct {
	Command    string
	Origin     CommandOriginData
	IsInternal bool
	Version    string
}

// Marshal reads or writes CommandRequest using its canonical wire layout.
func (x *CommandRequest) Marshal(io IO) {
	io.String(&x.Command)
	x.Origin.Marshal(io)
	io.Bool(&x.IsInternal)
	io.String(&x.Version)
}
