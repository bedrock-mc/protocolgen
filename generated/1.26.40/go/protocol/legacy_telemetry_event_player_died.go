// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventPlayerDied struct {
	InstigatorActorID    int32
	InstigatorMobVariant int32
	DamageSource         int32
	DiedInRaid           bool
}

func (*LegacyTelemetryEventPlayerDied) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPlayerDied using its canonical wire layout.
func (x *LegacyTelemetryEventPlayerDied) Marshal(io IO) {
	io.Varint32(&x.InstigatorActorID)
	io.Varint32(&x.InstigatorMobVariant)
	io.Varint32(&x.DamageSource)
	io.Bool(&x.DiedInRaid)
}
