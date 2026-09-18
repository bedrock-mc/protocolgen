// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// Camera is sent by the server to use an Education Edition camera on a player. It produces an image
// client-side.
type Camera struct {
	CameraID       int64
	TargetPlayerID int64
}

// ID ...
func (*Camera) ID() uint32 {
	return IDCamera
}

func (pk *Camera) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.CameraID)
	io.ActorUniqueID(&pk.TargetPlayerID)
}
