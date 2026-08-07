// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InitializeRegistryData struct {
	ClockData []WorldClockData
}

func (*InitializeRegistryData) isSyncWorldClocksData() {}

// Marshal reads or writes InitializeRegistryData using its canonical wire layout.
func (x *InitializeRegistryData) Marshal(io IO) {
	Slice(io, &x.ClockData)
}
