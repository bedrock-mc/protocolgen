// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetScoreboardIdentity struct {
	ScoreboardIdentityPacketType ScoreboardIdentityPacketType
	ScoreboardIdentityInfo       []ScoreboardIdentityPacketInfo
}

// Marshal reads or writes SetScoreboardIdentity using its canonical wire layout.
func (x *SetScoreboardIdentity) Marshal(io IO) {
	enumValue1 := uint8(x.ScoreboardIdentityPacketType)
	io.Uint8(&enumValue1)
	x.ScoreboardIdentityPacketType = ScoreboardIdentityPacketType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	if !io.Reading() && uint64(len(x.ScoreboardIdentityInfo)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ScoreboardIdentityInfo), "collection length overflows uint32")
		return
	}
	count2 := uint32(len(x.ScoreboardIdentityInfo))
	io.Varuint32(&count2)
	if io.Reading() {
		if uint64(count2) > uint64(^uint(0)>>1) {
			io.InvalidValue(count2, "collection length overflows int")
			return
		}
		x.ScoreboardIdentityInfo = make([]ScoreboardIdentityPacketInfo, int(count2))
	}
	for index3 := range x.ScoreboardIdentityInfo {
		x.ScoreboardIdentityInfo[index3].Marshal(io)
	}
}
