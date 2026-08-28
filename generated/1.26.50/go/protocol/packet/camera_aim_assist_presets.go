// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// CameraAimAssistPresets is sent by the server to the client to provide a list of categories and
// presets that can be used when sending a CameraAimAssist packet or a CameraInstruction including
// aim assist.
type CameraAimAssistPresets struct {
	// CameraAimAssistPresets is a list of categories which can be referenced by one of the Presets.
	CameraAimAssistPresets []protocol.CameraAimAssistCategoryDefinition
	// CameraAimAssistCategories is a list of presets which define a base for how aim assist should
	// behave
	CameraAimAssistCategories []protocol.CameraAimAssistPresetDefinition
	// Operation is the operation to perform with the presets. It is one of the constants above.
	Operation protocol.CameraAimAssistPresetOperation
}

// Marshal reads or writes CameraAimAssistPresets using its canonical wire layout.
func (x *CameraAimAssistPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.CameraAimAssistPresets)
	protocol.Slice(io, &x.CameraAimAssistCategories)
	protocol.IntegerFunc(&x.Operation, io.Uint8)
}

// ID returns the protocol ID for CameraAimAssistPresets.
func (*CameraAimAssistPresets) ID() uint32 { return IDCameraAimAssistPresets }
