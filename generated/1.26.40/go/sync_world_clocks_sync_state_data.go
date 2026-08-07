// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncWorldClocksSyncStateData struct {
	ClockData []SyncWorldClockStateData
}

func (SyncWorldClocksSyncStateData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncWorldClocksSyncStateData using its canonical wire layout.
func (x *SyncWorldClocksSyncStateData) Marshal(io IO) {
	FuncSlice(io, &x.ClockData, io.Varuint32, func(value *SyncWorldClockStateData) {
		value.Marshal(io)
	})
}
