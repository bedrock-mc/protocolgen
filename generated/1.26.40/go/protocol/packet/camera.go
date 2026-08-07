// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// Camera is sent by the server to use an Education Edition camera on a player. It produces an image
// client-side.
type Camera struct {
	CameraID       int64
	TargetPlayerID int64
}

// Marshal reads or writes Camera using its canonical wire layout.
func (x *Camera) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.CameraID)
	io.ActorUniqueID(&x.TargetPlayerID)
}

// ID returns the protocol ID for Camera.
func (*Camera) ID() uint32 { return IDCamera }
