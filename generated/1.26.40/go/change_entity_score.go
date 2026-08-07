// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ChangeEntityScore struct {
	Action        string
	ScoreboardId  ScoreboardId
	ObjectiveName string
	ScoreValue    int32
	ActorId       int64
}

func (ChangeEntityScore) isSetScoreScoreInfoItem() {}

// Marshal reads or writes ChangeEntityScore using its canonical wire layout.
func (x *ChangeEntityScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	io.ActorUniqueID(&x.ActorId)
}
