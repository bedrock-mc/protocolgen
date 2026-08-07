// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandOutput struct {
	OriginData CommandOriginData
	Output     CommandOutputData
}

// Marshal reads or writes CommandOutput using its canonical wire layout.
func (x *CommandOutput) Marshal(io IO) {
	x.OriginData.Marshal(io)
	x.Output.Marshal(io)
}
