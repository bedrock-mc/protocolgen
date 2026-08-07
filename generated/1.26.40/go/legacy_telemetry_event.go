// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEvent struct {
	TargetActorID ActorUniqueID
	EventType     LegacyTelemetryEventType
	UsePlayerID   bool
	EventData     LegacyTelemetryEventEventData
}
