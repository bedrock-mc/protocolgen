// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraSplineInstruction struct {
	TotalTime         float32
	Type              uint8
	Curve             []mgl32.Vec3
	ProgressKeyFrames []CameraProgressOption
	RotationOption    []CameraRotationOption
	SplineIdentifier  string
	LoadFromJSON      bool
}

// Marshal reads or writes CameraSplineInstruction using its canonical wire layout.
func (x *CameraSplineInstruction) Marshal(io IO) {
	io.Float32(&x.TotalTime)
	io.Uint8(&x.Type)
	FuncSlice(io, &x.Curve, io.Varuint32, io.Vec3)
	Slice(io, &x.ProgressKeyFrames)
	Slice(io, &x.RotationOption)
	io.String(&x.SplineIdentifier)
	io.Bool(&x.LoadFromJSON)
}
