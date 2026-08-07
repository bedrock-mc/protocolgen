// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
)

type PrimitiveShapeData struct {
	NetworkId             uint64
	ShapeType             Optional[ScriptModuleMinecraftScriptPrimitiveShapeType]
	Location              Optional[mgl32.Vec3]
	Scale                 Optional[float32]
	Rotation              Optional[mgl32.Vec3]
	TotalTimeLeft         Optional[float32]
	MaximumRenderDistance Optional[float32]
	Color                 Optional[color.RGBA]
	DimensionID           Optional[DimensionType]
	AttachedToEntityID    Optional[int64]
	ExtraShapeData        PrimitiveShapeDataExtraShapeData
}

// Marshal reads or writes PrimitiveShapeData using its canonical wire layout.
func (x *PrimitiveShapeData) Marshal(io IO) {
	io.Varuint64(&x.NetworkId)
	OptionalFunc(io, &x.ShapeType, func(value *ScriptModuleMinecraftScriptPrimitiveShapeType) {
		item := *value
		IntegerFunc(&item, io.Uint8)
		*value = item
	})
	OptionalFunc(io, &x.Location, io.Vec3)
	OptionalFunc(io, &x.Scale, io.Float32)
	OptionalFunc(io, &x.Rotation, io.Vec3)
	OptionalFunc(io, &x.TotalTimeLeft, io.Float32)
	OptionalFunc(io, &x.MaximumRenderDistance, io.Float32)
	OptionalFunc(io, &x.Color, io.RGBA)
	OptionalFunc(io, &x.DimensionID, func(value *DimensionType) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.AttachedToEntityID, io.ActorUniqueID)
	marshalPrimitiveShapeDataExtraShapeData(io, &x.ExtraShapeData)
}
