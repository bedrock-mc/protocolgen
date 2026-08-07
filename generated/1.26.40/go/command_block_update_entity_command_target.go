// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandBlockUpdateEntityCommandTarget struct {
	TargetRuntimeID uint64
}

func (CommandBlockUpdateEntityCommandTarget) isCommandBlockUpdateTarget() {}

// Marshal reads or writes CommandBlockUpdateEntityCommandTarget using its canonical wire layout.
func (x *CommandBlockUpdateEntityCommandTarget) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
}
