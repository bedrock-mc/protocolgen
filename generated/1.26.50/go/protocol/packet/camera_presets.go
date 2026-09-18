// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// CameraPresets gives the client a list of custom camera presets.
type CameraPresets struct {
	Presets []protocol.CameraPreset
}

// ID returns the protocol ID for CameraPresets.
func (*CameraPresets) ID() uint32 { return IDCameraPresets }

// Marshal reads or writes CameraPresets using its canonical wire layout.
func (pk *CameraPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Presets)
}
