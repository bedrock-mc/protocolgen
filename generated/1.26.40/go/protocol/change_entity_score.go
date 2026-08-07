// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChangeEntityScore struct {
	Action        string
	ScoreboardID  ScoreboardID
	ObjectiveName string
	ScoreValue    int32
	ActorID       int64
}

func (*ChangeEntityScore) isSetScoreInfoItem() {}

// Marshal reads or writes ChangeEntityScore using its canonical wire layout.
func (x *ChangeEntityScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardID.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	io.ActorUniqueID(&x.ActorID)
}
