// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraPresetsData struct {
	Presets []CameraPreset
}

// Marshal reads or writes CameraPresetsData using its canonical wire layout.
func (x *CameraPresetsData) Marshal(io IO) {
	FuncSlice(io, &x.Presets, io.Varuint32, func(value *CameraPreset) {
		value.Marshal(io)
	})
}
