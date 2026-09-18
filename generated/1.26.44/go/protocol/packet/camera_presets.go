// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// CameraPresets gives the client a list of custom camera presets.
type CameraPresets struct {
	// Presets is a list of camera presets that can be used by other cameras. The order of this list is important
	// because the index of presets is used as a pointer in the CameraInstruction packet.
	Presets []protocol.CameraPreset
}

// ID returns the protocol ID for CameraPresets.
func (*CameraPresets) ID() uint32 { return IDCameraPresets }

// Marshal reads or writes CameraPresets using its canonical wire layout.
func (pk *CameraPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Presets)
}
