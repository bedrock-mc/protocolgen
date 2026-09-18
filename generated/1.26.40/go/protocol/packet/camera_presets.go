// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// CameraPresets gives the client a list of custom camera presets.
type CameraPresets struct {
	Presets []protocol.CameraPreset
}

// ID ...
func (*CameraPresets) ID() uint32 {
	return IDCameraPresets
}

func (pk *CameraPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Presets)
}
