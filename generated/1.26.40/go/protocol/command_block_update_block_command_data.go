// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandBlockUpdateBlockCommandData struct {
	BlockPosition    BlockPos
	CommandBlockMode uint32
	RedstoneMode     bool
	IsConditional    bool
}

func (*CommandBlockUpdateBlockCommandData) isCommandBlockUpdateTarget() {}

// Marshal reads or writes CommandBlockUpdateBlockCommandData using its canonical wire layout.
func (x *CommandBlockUpdateBlockCommandData) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.Varuint32(&x.CommandBlockMode)
	io.Bool(&x.RedstoneMode)
	io.Bool(&x.IsConditional)
}
