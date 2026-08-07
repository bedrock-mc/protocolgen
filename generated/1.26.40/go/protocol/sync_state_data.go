// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncStateData struct {
	ClockData []SyncWorldClockStateData
}

func (*SyncStateData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncStateData using its canonical wire layout.
func (x *SyncStateData) Marshal(io IO) {
	Slice(io, &x.ClockData)
}
