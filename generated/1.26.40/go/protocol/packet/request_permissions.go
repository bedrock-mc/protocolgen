// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RequestPermissions struct {
	TargetPlayerIdSRawID  int64
	PlayerPermissionLevel int32
	CustomPermissionFlags uint16
}

// Marshal reads or writes RequestPermissions using its canonical wire layout.
func (x *RequestPermissions) Marshal(io protocol.IO) {
	io.Int64(&x.TargetPlayerIdSRawID)
	io.Varint32(&x.PlayerPermissionLevel)
	io.Uint16(&x.CustomPermissionFlags)
}
