// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventMobBorn struct {
	BornBabyEntityType    int32
	BornBabyEntityVariant int32
	BornBabyColor         uint8
}

func (LegacyTelemetryEventMobBorn) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventMobBorn using its canonical wire layout.
func (x *LegacyTelemetryEventMobBorn) Marshal(io IO) {
	io.Varint32(&x.BornBabyEntityType)
	io.Varint32(&x.BornBabyEntityVariant)
	io.Uint8(&x.BornBabyColor)
}
