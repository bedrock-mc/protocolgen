// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraSpline struct {
	CameraDataSplines []CameraSplineDefinition
}

// Marshal reads or writes CameraSpline using its canonical wire layout.
func (x *CameraSpline) Marshal(io IO) {
	FuncSlice(io, &x.CameraDataSplines, io.Varuint32, func(value *CameraSplineDefinition) {
		value.Marshal(io)
	})
}
