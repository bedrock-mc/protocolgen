// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventMobKilled struct {
	InstigatorActorID         int64
	TargetActorID             int64
	InstigatorSChildActorType ActorType
	DamageSource              int32
	TradeTier                 int32
	TraderName                string
}

func (LegacyTelemetryEventMobKilled) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventMobKilled using its canonical wire layout.
func (x *LegacyTelemetryEventMobKilled) Marshal(io IO) {
	io.Varint64(&x.InstigatorActorID)
	io.Varint64(&x.TargetActorID)
	IntegerFunc(&x.InstigatorSChildActorType, io.Varint32)
	io.Varint32(&x.DamageSource)
	io.Varint32(&x.TradeTier)
	io.String(&x.TraderName)
}
