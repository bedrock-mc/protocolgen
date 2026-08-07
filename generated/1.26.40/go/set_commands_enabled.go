// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetCommandsEnabled struct {
	CommandsEnabled bool
}

// Marshal reads or writes SetCommandsEnabled using its canonical wire layout.
func (x *SetCommandsEnabled) Marshal(io IO) {
	io.Bool(&x.CommandsEnabled)
}
