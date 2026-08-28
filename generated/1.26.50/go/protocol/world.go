// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// DimensionDefinition contains information specifying dimension-specific properties, used for
// data-driven dimensions. These include the range (the height min/max), generator variant, and
// more.
type DimensionDefinition struct {
	MinimumY      int32
	HeightRange   int32
	GeneratorType GeneratorType
	// DimensionType is the numeric identifier of the dimension. This cannot override a vanilla
	// dimension (0-2), but custom dimensions should start from 1000 like vanilla.
	DimensionType DimensionType
	// PackID is the UUID of the behaviour pack which has added the dimension.
	PackID       uuid.UUID
	DefaultBiome string
}

// Marshal reads or writes DimensionDefinition using its canonical wire layout.
func (x *DimensionDefinition) Marshal(io IO) {
	io.Varint32(&x.MinimumY)
	io.Varint32(&x.HeightRange)
	IntegerFunc(&x.GeneratorType, io.Varint32)
	x.DimensionType.Marshal(io)
	io.UUID(&x.PackID)
	io.StringLimits(&x.DefaultBiome, 0, 256)
}

type WorldPosition struct {
	Position      mgl32.Vec3
	DimensionType DimensionType
}

// Marshal reads or writes WorldPosition using its canonical wire layout.
func (x *WorldPosition) Marshal(io IO) {
	io.Vec3(&x.Position)
	x.DimensionType.Marshal(io)
}
