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
	protocol.Slice(io, &x.CameraAimAssistPresets)
	protocol.Slice(io, &x.CameraAimAssistCategories)
	protocol.IntegerFunc(&x.Operation, io.Uint8)
}

// ID returns the protocol ID for CameraAimAssistPresets.
func (*CameraAimAssistPresets) ID() uint32 { return IDCameraAimAssistPresets }
