// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CameraPresets struct {
	CameraPresets protocol.CameraPresetsData
}

// Marshal reads or writes CameraPresets using its canonical wire layout.
func (x *CameraPresets) Marshal(io protocol.IO) {
	x.CameraPresets.Marshal(io)
}
