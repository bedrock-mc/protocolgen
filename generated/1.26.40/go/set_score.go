// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetScore struct {
	ScoreInfo []SetScoreScoreInfoItem
}

// Marshal reads or writes SetScore using its canonical wire layout.
func (x *SetScore) Marshal(io IO) {
	FuncSlice(io, &x.ScoreInfo, io.Varuint32, func(value *SetScoreScoreInfoItem) {
		marshalSetScoreScoreInfoItem(io, value)
	})
}
