// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CameraAimAssistPresets struct {
	CameraAimAssistPresets    []protocol.CameraAimAssistCategoryDefinition
	CameraAimAssistCategories []protocol.CameraAimAssistPresetDefinition
	Operation                 protocol.CameraAimAssistPresetsPacketOperation
}

// Marshal reads or writes CameraAimAssistPresets using its canonical wire layout.
func (x *CameraAimAssistPresets) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.CameraAimAssistPresets, io.Varuint32, func(value *protocol.CameraAimAssistCategoryDefinition) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.CameraAimAssistCategories, io.Varuint32, func(value *protocol.CameraAimAssistPresetDefinition) {
		value.Marshal(io)
	})
	protocol.IntegerFunc(&x.Operation, io.Uint8)
}
