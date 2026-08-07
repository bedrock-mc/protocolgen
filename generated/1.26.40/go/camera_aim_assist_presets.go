// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistPresets struct {
	CameraAimAssistPresets    []CameraAimAssistCategoryDefinition
	CameraAimAssistCategories []CameraAimAssistPresetDefinition
	Operation                 CameraAimAssistPresetsPacketOperation
}

// Marshal reads or writes CameraAimAssistPresets using its canonical wire layout.
func (x *CameraAimAssistPresets) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.CameraAimAssistPresets)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CameraAimAssistPresets), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.CameraAimAssistPresets))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.CameraAimAssistPresets = make([]CameraAimAssistCategoryDefinition, int(count1))
	}
	for index2 := range x.CameraAimAssistPresets {
		x.CameraAimAssistPresets[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.CameraAimAssistCategories)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CameraAimAssistCategories), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.CameraAimAssistCategories))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.CameraAimAssistCategories = make([]CameraAimAssistPresetDefinition, int(count3))
	}
	for index4 := range x.CameraAimAssistCategories {
		x.CameraAimAssistCategories[index4].Marshal(io)
	}
	enumValue5 := uint8(x.Operation)
	io.Uint8(&enumValue5)
	x.Operation = CameraAimAssistPresetsPacketOperation(enumValue5)
	switch int64(enumValue5) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue5, "unknown enum value")
	}
}
