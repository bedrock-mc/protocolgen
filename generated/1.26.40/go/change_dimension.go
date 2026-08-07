// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type ChangeDimension struct {
	DimensionID     DimensionType
	Position        mgl32.Vec3
	Respawn         bool
	LoadingScreenId Optional[uint32]
}
