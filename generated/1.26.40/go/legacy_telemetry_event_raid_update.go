// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventRaidUpdate struct {
	CurrentWave int32
	TotalWaves  int32
	Success     bool
}

func (LegacyTelemetryEventRaidUpdate) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventRaidUpdate using its canonical wire layout.
func (x *LegacyTelemetryEventRaidUpdate) Marshal(io IO) {
	io.Varint32(&x.CurrentWave)
	io.Varint32(&x.TotalWaves)
	io.Bool(&x.Success)
}
