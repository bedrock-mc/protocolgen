// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CameraSpline struct {
	CameraDataSplines []protocol.CameraSplineDefinition
}

// Marshal reads or writes CameraSpline using its canonical wire layout.
func (x *CameraSpline) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.CameraDataSplines, io.Varuint32, func(value *protocol.CameraSplineDefinition) {
		value.Marshal(io)
	})
}
