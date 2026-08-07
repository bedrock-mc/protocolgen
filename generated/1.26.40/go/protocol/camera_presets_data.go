// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraPresetsData struct {
	Presets []CameraPreset
}

// Marshal reads or writes CameraPresetsData using its canonical wire layout.
func (x *CameraPresetsData) Marshal(io IO) {
	Slice(io, &x.Presets)
}
