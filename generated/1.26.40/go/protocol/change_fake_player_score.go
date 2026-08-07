// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChangeFakePlayerScore struct {
	Action         string
	ScoreboardID   ScoreboardID
	ObjectiveName  string
	ScoreValue     int32
	FakePlayerName string
}

func (*ChangeFakePlayerScore) isSetScoreInfoItem() {}

// Marshal reads or writes ChangeFakePlayerScore using its canonical wire layout.
func (x *ChangeFakePlayerScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardID.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	io.String(&x.FakePlayerName)
}
