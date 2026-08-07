// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetScoreboardIdentity struct {
	ScoreboardIdentityPacketType ScoreboardIdentityPacketType
	ScoreboardIdentityInfo       []ScoreboardIdentityPacketInfo
}

// Marshal reads or writes SetScoreboardIdentity using its canonical wire layout.
func (x *SetScoreboardIdentity) Marshal(io IO) {
	IntegerFunc(&x.ScoreboardIdentityPacketType, io.Uint8)
	FuncSlice(io, &x.ScoreboardIdentityInfo, io.Varuint32, func(value *ScoreboardIdentityPacketInfo) {
		value.Marshal(io)
	})
}
