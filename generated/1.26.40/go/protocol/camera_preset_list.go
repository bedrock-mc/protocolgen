// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraPresetList struct {
	Presets []CameraPreset
}

// Marshal reads or writes CameraPresetList using its canonical wire layout.
func (x *CameraPresetList) Marshal(io IO) {
	Slice(io, &x.Presets)
}
