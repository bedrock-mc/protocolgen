// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
)

type ClientboundDebugRendererDebugMarkerData struct {
	Text     string
	Position mgl32.Vec3
	Color    color.RGBA
	Duration uint64
}

// Marshal reads or writes ClientboundDebugRendererDebugMarkerData using its canonical wire layout.
func (x *ClientboundDebugRendererDebugMarkerData) Marshal(io IO) {
	io.String(&x.Text)
	io.Vec3(&x.Position)
	io.RGBA(&x.Color)
	io.Uint64(&x.Duration)
}
