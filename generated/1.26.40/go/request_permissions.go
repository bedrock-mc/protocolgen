// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RequestPermissions struct {
	TargetPlayerIdSRawID  int64
	PlayerPermissionLevel int32
	CustomPermissionFlags uint16
}

// Marshal reads or writes RequestPermissions using its canonical wire layout.
func (x *RequestPermissions) Marshal(io IO) {
	io.Int64(&x.TargetPlayerIdSRawID)
	io.Varint32(&x.PlayerPermissionLevel)
	io.Uint16(&x.CustomPermissionFlags)
}
