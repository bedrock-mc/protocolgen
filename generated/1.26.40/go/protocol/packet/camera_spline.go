// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CameraSpline is sent by the server to define camera spline paths.
type CameraSpline struct {
	// Splines is a list of camera spline definitions.
	CameraDataSplines []protocol.CameraSplineDefinition
}

// Marshal reads or writes CameraSpline using its canonical wire layout.
func (x *CameraSpline) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.CameraDataSplines)
}

// ID returns the protocol ID for CameraSpline.
func (*CameraSpline) ID() uint32 { return IDCameraSpline }
