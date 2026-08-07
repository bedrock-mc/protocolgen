// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChangePlayerScore struct {
	Action         string
	ScoreboardId   ScoreboardId
	ObjectiveName  string
	ScoreValue     int32
	PlayerUniqueId PlayerScoreboardId
}

func (*ChangePlayerScore) isSetScoreScoreInfoItem() {}

// Marshal reads or writes ChangePlayerScore using its canonical wire layout.
func (x *ChangePlayerScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	x.PlayerUniqueId.Marshal(io)
}
