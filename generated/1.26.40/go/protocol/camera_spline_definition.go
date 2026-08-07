// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraSplineDefinition struct {
	Name              string
	TotalTime         float32
	SplineType        string
	ControlPoints     []CameraSplineControlPoint
	ProgressKeyFrames []CameraSplineProgressKeyFrame
	RotationKeyFrames []CameraSplineRotationKeyFrame
}

// Marshal reads or writes CameraSplineDefinition using its canonical wire layout.
func (x *CameraSplineDefinition) Marshal(io IO) {
	io.String(&x.Name)
	io.Float32(&x.TotalTime)
	io.String(&x.SplineType)
	Slice(io, &x.ControlPoints)
	Slice(io, &x.ProgressKeyFrames)
	Slice(io, &x.RotationKeyFrames)
}
