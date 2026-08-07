// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventActorDefinition struct {
	EventName string
}

func (LegacyTelemetryEventActorDefinition) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventActorDefinition using its canonical wire layout.
func (x *LegacyTelemetryEventActorDefinition) Marshal(io IO) {
	io.String(&x.EventName)
}
