// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEvent struct {
	TargetActorID ActorUniqueID
	EventType     LegacyTelemetryEventType
	UsePlayerID   bool
	EventData     LegacyTelemetryEventEventData
}

// Marshal reads or writes LegacyTelemetryEvent using its canonical wire layout.
func (x *LegacyTelemetryEvent) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	enumValue1 := int32(x.EventType)
	io.Varint32(&enumValue1)
	x.EventType = LegacyTelemetryEventType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.UsePlayerID)
	marshalLegacyTelemetryEventEventData(io, &x.EventData)
}
