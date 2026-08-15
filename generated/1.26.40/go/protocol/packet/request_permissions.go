// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// RequestPermissions is a packet sent from the client to the server to request permissions that the
// client does not currently have. It can only be sent by operators and host in vanilla Minecraft.
type RequestPermissions struct {
	// TargetPlayerIDSRawID is the unique ID of the player. The unique ID is unique for the entire world
	// and is often used in packets. Most servers send an EntityUniqueID equal to the EntityRuntimeID.
	TargetPlayerIDSRawID int64
	// PlayerPermissionLevel is the current permission level of the player. This is one of the constants
	// that may be found in the AdventureSettings packet.
	PlayerPermissionLevel int32
	// CustomPermissionFlags contains the requested permission flags.
	CustomPermissionFlags uint16
}

// Marshal reads or writes RequestPermissions using its canonical wire layout.
func (x *RequestPermissions) Marshal(io protocol.IO) {
	io.Int64(&x.TargetPlayerIDSRawID)
	io.Varint32(&x.PlayerPermissionLevel)
	io.Uint16(&x.CustomPermissionFlags)
	protocol.Minimum(io, &x.CustomPermissionFlags, 0)
}

// ID returns the protocol ID for RequestPermissions.
func (*RequestPermissions) ID() uint32 { return IDRequestPermissions }
