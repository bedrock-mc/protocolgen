// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerScoreboardID struct {
	PlayerUniqueID int64
}

// Marshal reads or writes PlayerScoreboardID using its canonical wire layout.
func (x *PlayerScoreboardID) Marshal(io IO) {
	io.Varint64(&x.PlayerUniqueID)
}
