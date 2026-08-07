// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CameraPresets gives the client a list of custom camera presets.
type CameraPresets struct {
	CameraPresets protocol.CameraPresetList
}

// Marshal reads or writes CameraPresets using its canonical wire layout.
func (x *CameraPresets) Marshal(io protocol.IO) {
	x.CameraPresets.Marshal(io)
}

// ID returns the protocol ID for CameraPresets.
func (*CameraPresets) ID() uint32 { return IDCameraPresets }
