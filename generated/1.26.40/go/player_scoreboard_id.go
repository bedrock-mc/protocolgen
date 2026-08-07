// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerScoreboardId struct {
	PlayerUniqueId int64
}

// Marshal reads or writes PlayerScoreboardId using its canonical wire layout.
func (x *PlayerScoreboardId) Marshal(io IO) {
	io.Varint64(&x.PlayerUniqueId)
}
