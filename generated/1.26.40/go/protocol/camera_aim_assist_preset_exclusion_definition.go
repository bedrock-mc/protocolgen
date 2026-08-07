// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraAimAssistPresetExclusionDefinition struct {
	Blocks             []string
	Entities           []string
	BlockTags          []string
	EntityTypeFamilies []string
}

// Marshal reads or writes CameraAimAssistPresetExclusionDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetExclusionDefinition) Marshal(io IO) {
	FuncSlice(io, &x.Blocks, io.Varuint32, io.String)
	FuncSlice(io, &x.Entities, io.Varuint32, io.String)
	FuncSlice(io, &x.BlockTags, io.Varuint32, io.String)
	FuncSlice(io, &x.EntityTypeFamilies, io.Varuint32, io.String)
}
