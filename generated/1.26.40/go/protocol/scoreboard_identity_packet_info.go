// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ScoreboardIdentityPacketInfo struct {
	ScoreboardID   ScoreboardID
	PlayerUniqueID Optional[int64]
}

// Marshal reads or writes ScoreboardIdentityPacketInfo using its canonical wire layout.
func (x *ScoreboardIdentityPacketInfo) Marshal(io IO) {
	x.ScoreboardID.Marshal(io)
	OptionalFunc(io, &x.PlayerUniqueID, io.Varint64)
}
