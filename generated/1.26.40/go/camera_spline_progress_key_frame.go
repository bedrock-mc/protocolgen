// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraSplineProgressKeyFrame struct {
	Progress float32
	Time     float32
	Easing   Optional[string]
}

// Marshal reads or writes CameraSplineProgressKeyFrame using its canonical wire layout.
func (x *CameraSplineProgressKeyFrame) Marshal(io IO) {
	io.Float32(&x.Progress)
	io.Float32(&x.Time)
	OptionalFunc(io, &x.Easing, io.String)
}
