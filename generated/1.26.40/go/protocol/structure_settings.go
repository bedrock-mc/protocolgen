// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type StructureSettings struct {
	StructurePaletteName                            string
	ShouldIgnoreEntities                            bool
	ShouldIgnoreBlocks                              bool
	ShouldAllowNonTickingPlayerAndTickingAreaChunks bool
	StructureSize                                   BlockPos
	StructureOffset                                 BlockPos
	LastEditPlayer                                  int64
	Rotation                                        Rotation
	Mirror                                          Mirror
	AnimationMode                                   AnimationMode
	AnimationSeconds                                float32
	IntegrityValue                                  float32
	IntegritySeed                                   uint32
	RotationPivot                                   mgl32.Vec3
}

// Marshal reads or writes StructureSettings using its canonical wire layout.
func (x *StructureSettings) Marshal(io IO) {
	io.String(&x.StructurePaletteName)
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
