// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RemoveScore struct {
	Action        string
	ScoreboardId  ScoreboardId
	ObjectiveName Optional[string]
}

func (RemoveScore) isSetScoreScoreInfoItem() {}

// Marshal reads or writes RemoveScore using its canonical wire layout.
func (x *RemoveScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	OptionalFunc(io, &x.ObjectiveName, io.String)
}
