// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ScoreboardId struct {
	ScoreboardId int64
}

// Marshal reads or writes ScoreboardId using its canonical wire layout.
func (x *ScoreboardId) Marshal(io IO) {
	io.Varint64(&x.ScoreboardId)
}
