// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncWorldClocksSyncStateData struct {
	ClockData []SyncWorldClockStateData
}

func (*SyncWorldClocksSyncStateData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncWorldClocksSyncStateData using its canonical wire layout.
func (x *SyncWorldClocksSyncStateData) Marshal(io IO) {
	Slice(io, &x.ClockData)
}
