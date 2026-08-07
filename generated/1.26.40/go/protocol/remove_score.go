// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RemoveScore struct {
	Action        string
	ScoreboardID  ScoreboardID
	ObjectiveName Optional[string]
}

func (*RemoveScore) isSetScoreInfoItem() {}

// Marshal reads or writes RemoveScore using its canonical wire layout.
func (x *RemoveScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardID.Marshal(io)
	OptionalFunc(io, &x.ObjectiveName, io.String)
}
