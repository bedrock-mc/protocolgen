// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraAimAssistCategoryPriorities struct {
	Entities           []OrderedEntry[string, int32]
	Blocks             []OrderedEntry[string, int32]
	BlockTags          []OrderedEntry[string, int32]
	EntityTypeFamilies []OrderedEntry[string, int32]
	EntityDefault      Optional[int32]
	BlockDefault       Optional[int32]
}

// Marshal reads or writes CameraAimAssistCategoryPriorities using its canonical wire layout.
func (x *CameraAimAssistCategoryPriorities) Marshal(io IO) {
	OrderedMap(io, &x.Entities, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.Blocks, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.BlockTags, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.EntityTypeFamilies, io.Varuint32, io.String, io.Int32)
	OptionalFunc(io, &x.EntityDefault, io.Int32)
	OptionalFunc(io, &x.BlockDefault, io.Int32)
}
