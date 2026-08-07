// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventBossKilled struct {
	BossActorID int64
	PartySize   int32
	BossType    int32
}

func (LegacyTelemetryEventBossKilled) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventBossKilled using its canonical wire layout.
func (x *LegacyTelemetryEventBossKilled) Marshal(io IO) {
	io.Varint64(&x.BossActorID)
	io.Varint32(&x.PartySize)
	io.Varint32(&x.BossType)
}
