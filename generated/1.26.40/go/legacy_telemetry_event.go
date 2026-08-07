// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEvent struct {
	TargetActorID int64
	EventType     LegacyTelemetryEventType
	UsePlayerID   bool
	EventData     LegacyTelemetryEventEventData
}

// Marshal reads or writes LegacyTelemetryEvent using its canonical wire layout.
func (x *LegacyTelemetryEvent) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetActorID)
	IntegerFunc(&x.EventType, io.Varint32)
	io.Bool(&x.UsePlayerID)
	marshalLegacyTelemetryEventEventData(io, &x.EventData)
}
