// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// PositionTrackingDBServerBroadcast is sent by the server in response to the
// PositionTrackingDBClientRequest packet. This packet is, as of 1.16, currently only used for
// lodestones. The server maintains a database with tracking IDs and their position and dimension.
// The client will request these tracking IDs, (NBT tag set on the lodestone compass with the
// tracking ID?) and the server will respond with the status of those tracking IDs. What is actually
// done with the data sent depends on what the client chooses to do with it. For the lodestone
// compass, it is used to make the compass point towards lodestones and to make it spin if the
// lodestone at a position is no longer there.
type PositionTrackingDBServerBroadcast struct {
	Action               protocol.PositionTrackingDBServerBroadcastAction
	IDValue              protocol.PositionTrackingID
	PositionTrackingData []byte
}

// Marshal reads or writes PositionTrackingDBServerBroadcast using its canonical wire layout.
func (x *PositionTrackingDBServerBroadcast) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Action, io.Uint8)
	x.IDValue.Marshal(io)
	io.NBT(&x.PositionTrackingData, protocol.NBTNetwork)
}

// ID returns the protocol ID for PositionTrackingDBServerBroadcast.
func (*PositionTrackingDBServerBroadcast) ID() uint32 { return IDPositionTrackingDBServerBroadcast }
