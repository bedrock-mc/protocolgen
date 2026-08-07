// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraPresets struct {
	CameraPresets CameraPresetsData
}

// Marshal reads or writes CameraPresets using its canonical wire layout.
func (x *CameraPresets) Marshal(io IO) {
	x.CameraPresets.Marshal(io)
}
