// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraSpline struct {
	CameraDataSplines []CameraSplineDefinition
}

// Marshal reads or writes CameraSpline using its canonical wire layout.
func (x *CameraSpline) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.CameraDataSplines)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CameraDataSplines), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.CameraDataSplines))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.CameraDataSplines = make([]CameraSplineDefinition, int(count1))
	}
	for index2 := range x.CameraDataSplines {
		x.CameraDataSplines[index2].Marshal(io)
	}
}
