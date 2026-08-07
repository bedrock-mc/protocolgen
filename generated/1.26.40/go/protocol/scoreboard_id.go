// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ScoreboardID struct {
	ScoreboardID int64
}

// Marshal reads or writes ScoreboardID using its canonical wire layout.
func (x *ScoreboardID) Marshal(io IO) {
	io.Varint64(&x.ScoreboardID)
}
