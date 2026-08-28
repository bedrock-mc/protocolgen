// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// CameraInstruction gives a custom camera specific instructions to operate.
type CameraInstruction struct {
	CameraInstruction protocol.CameraInstructionData
}

// Marshal reads or writes CameraInstruction using its canonical wire layout.
func (x *CameraInstruction) Marshal(io protocol.IO) {
	x.CameraInstruction.Marshal(io)
}

// ID returns the protocol ID for CameraInstruction.
func (*CameraInstruction) ID() uint32 { return IDCameraInstruction }
