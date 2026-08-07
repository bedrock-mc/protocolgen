// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetScore struct {
	ScoreInfo []SetScoreScoreInfoItem
}

// Marshal reads or writes SetScore using its canonical wire layout.
func (x *SetScore) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ScoreInfo)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ScoreInfo), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ScoreInfo))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ScoreInfo = make([]SetScoreScoreInfoItem, int(count1))
	}
	for index2 := range x.ScoreInfo {
		marshalSetScoreScoreInfoItem(io, &x.ScoreInfo[index2])
	}
}
