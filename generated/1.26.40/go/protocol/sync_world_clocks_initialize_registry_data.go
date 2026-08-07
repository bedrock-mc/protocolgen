// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncWorldClocksInitializeRegistryData struct {
	ClockData []WorldClockData
}

func (SyncWorldClocksInitializeRegistryData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncWorldClocksInitializeRegistryData using its canonical wire layout.
func (x *SyncWorldClocksInitializeRegistryData) Marshal(io IO) {
	FuncSlice(io, &x.ClockData, io.Varuint32, func(value *WorldClockData) {
		value.Marshal(io)
	})
}
