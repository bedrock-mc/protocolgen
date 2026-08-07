// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type StructureEditorData struct {
	StructureName         BedrockSafetyRedactableString
	DataField             string
	ShouldIncludePlayers  bool
	ShouldShowBoundingBox bool
	StructureBlockType    StructureBlockType
	StructureSettings     StructureSettings
	RedstoneSaveMode      StructureRedstoneSaveMode
}

// Marshal reads or writes StructureEditorData using its canonical wire layout.
func (x *StructureEditorData) Marshal(io IO) {
	x.StructureName.Marshal(io)
	io.String(&x.DataField)
	io.Bool(&x.ShouldIncludePlayers)
	io.Bool(&x.ShouldShowBoundingBox)
	IntegerFunc(&x.StructureBlockType, io.Varint32)
	x.StructureSettings.Marshal(io)
	IntegerFunc(&x.RedstoneSaveMode, io.Uint8)
}
