// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistPresets struct {
	CameraAimAssistPresets    []CameraAimAssistCategoryDefinition
	CameraAimAssistCategories []CameraAimAssistPresetDefinition
	Operation                 CameraAimAssistPresetsPacketOperation
}

// Marshal reads or writes CameraAimAssistPresets using its canonical wire layout.
func (x *CameraAimAssistPresets) Marshal(io IO) {
	FuncSlice(io, &x.CameraAimAssistPresets, io.Varuint32, func(value *CameraAimAssistCategoryDefinition) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.CameraAimAssistCategories, io.Varuint32, func(value *CameraAimAssistPresetDefinition) {
		value.Marshal(io)
	})
	IntegerFunc(&x.Operation, io.Uint8)
}
