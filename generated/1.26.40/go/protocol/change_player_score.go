// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChangePlayerScore struct {
	Action         string
	ScoreboardID   ScoreboardID
	ObjectiveName  string
	ScoreValue     int32
	PlayerUniqueID PlayerScoreboardID
}

func (*ChangePlayerScore) isSetScoreInfoItem() {}

// Marshal reads or writes ChangePlayerScore using its canonical wire layout.
func (x *ChangePlayerScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardID.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	x.PlayerUniqueID.Marshal(io)
}
