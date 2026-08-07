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
	FuncSlice(io, &x.ControlPoints, io.Varuint32, func(value *CameraSplineControlPoint) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ProgressKeyFrames, io.Varuint32, func(value *CameraSplineProgressKeyFrame) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.RotationKeyFrames, io.Varuint32, func(value *CameraSplineRotationKeyFrame) {
		value.Marshal(io)
	})
}
