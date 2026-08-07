// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ScoreboardIdentityPacketInfo struct {
	ScoreboardId   ScoreboardId
	PlayerUniqueId Optional[int64]
}

// Marshal reads or writes ScoreboardIdentityPacketInfo using its canonical wire layout.
func (x *ScoreboardIdentityPacketInfo) Marshal(io IO) {
	x.ScoreboardId.Marshal(io)
	OptionalFunc(io, &x.PlayerUniqueId, io.Varint64)
}
