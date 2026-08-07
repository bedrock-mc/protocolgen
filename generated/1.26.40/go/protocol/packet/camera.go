// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Camera struct {
	CameraID       int64
	TargetPlayerID int64
}

// Marshal reads or writes Camera using its canonical wire layout.
func (x *Camera) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.CameraID)
	io.ActorUniqueID(&x.TargetPlayerID)
}
