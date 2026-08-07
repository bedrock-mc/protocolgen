// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraInstructionOptionsSplineInstruction struct {
	TotalTime         float32
	Type              uint8
	Curve             []mgl32.Vec3
	ProgressKeyFrames []CameraInstructionOptionsSplineInstructionSplineProgressOption
	RotationOption    []CameraInstructionOptionsSplineInstructionSplineRotationOption
	SplineIdentifier  string
	LoadFromJson      bool
}

// Marshal reads or writes CameraInstructionOptionsSplineInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstruction) Marshal(io IO) {
	io.Float32(&x.TotalTime)
	io.Uint8(&x.Type)
	FuncSlice(io, &x.Curve, io.Varuint32, io.Vec3)
	FuncSlice(io, &x.ProgressKeyFrames, io.Varuint32, func(value *CameraInstructionOptionsSplineInstructionSplineProgressOption) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.RotationOption, io.Varuint32, func(value *CameraInstructionOptionsSplineInstructionSplineRotationOption) {
		value.Marshal(io)
	})
	io.String(&x.SplineIdentifier)
	io.Bool(&x.LoadFromJson)
}
