// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type StructureBlockType int32

const (
	StructureBlockTypeData    StructureBlockType = 0
	StructureBlockTypeSave    StructureBlockType = 1
	StructureBlockTypeLoad    StructureBlockType = 2
	StructureBlockTypeCorner  StructureBlockType = 3
	StructureBlockTypeInvalid StructureBlockType = 4
	StructureBlockTypeExport  StructureBlockType = 5
)

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

type StructureRedstoneSaveMode uint8

const (
	StructureRedstoneSaveModeSavesToMemory StructureRedstoneSaveMode = 0
	StructureRedstoneSaveModeSavesToDisk   StructureRedstoneSaveMode = 1
)

// StructureSettings is a struct holding settings of a structure block. Its fields may be changed
// using the in-game UI on the client-side.
type StructureSettings struct {
	StructurePaletteName                            string
	ShouldIgnoreEntities                            bool
	ShouldIgnoreBlocks                              bool
	ShouldAllowNonTickingPlayerAndTickingAreaChunks bool
	StructureSize                                   BlockPos
	StructureOffset                                 BlockPos
	LastEditPlayer                                  int64
	// Rotation is the rotation that the structure block should obtain. See the constants above for
	// available options.
	Rotation Rotation
	// Mirror specifies the way the structure should be mirrored. It is either no mirror at all, mirror
	// on the x/z axis or both.
	Mirror Mirror
	// AnimationMode ...
	AnimationMode    AnimationMode
	AnimationSeconds float32
	IntegrityValue   float32
	IntegritySeed    uint32
	RotationPivot    mgl32.Vec3
}

// Marshal reads or writes StructureSettings using its canonical wire layout.
func (x *StructureSettings) Marshal(io IO) {
	io.StringLimits(&x.StructurePaletteName, 0, 256)
	io.Bool(&x.ShouldIgnoreEntities)
	io.Bool(&x.ShouldIgnoreBlocks)
	io.Bool(&x.ShouldAllowNonTickingPlayerAndTickingAreaChunks)
	x.StructureSize.Marshal(io)
	x.StructureOffset.Marshal(io)
	io.ActorUniqueID(&x.LastEditPlayer)
	IntegerFunc(&x.Rotation, io.Uint8)
	IntegerFunc(&x.Mirror, io.Uint8)
	IntegerFunc(&x.AnimationMode, io.Uint8)
	io.Float32(&x.AnimationSeconds)
	io.Float32(&x.IntegrityValue)
	io.Uint32(&x.IntegritySeed)
	io.Vec3(&x.RotationPivot)
}

type StructureTemplateRequestOperation uint8

const (
	StructureTemplateRequestOperationNone                StructureTemplateRequestOperation = 0
	StructureTemplateRequestOperationExportFromSaveMode  StructureTemplateRequestOperation = 1
	StructureTemplateRequestOperationExportFromLoadMode  StructureTemplateRequestOperation = 2
	StructureTemplateRequestOperationQuerySavedStructure StructureTemplateRequestOperation = 3
)

type StructureTemplateResponseType uint8

const (
	StructureTemplateResponseTypeNone   StructureTemplateResponseType = 0
	StructureTemplateResponseTypeExport StructureTemplateResponseType = 1
	StructureTemplateResponseTypeQuery  StructureTemplateResponseType = 2
)
