// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsChainedSubcommandRelationship struct {
	SubCommandFirstValue  uint32
	SubCommandSecondValue uint32
}

// Marshal reads or writes AvailableCommandsChainedSubcommandRelationship using its canonical wire layout.
func (x *AvailableCommandsChainedSubcommandRelationship) Marshal(io IO) {
	io.Varuint32(&x.SubCommandFirstValue)
	io.Varuint32(&x.SubCommandSecondValue)
}
