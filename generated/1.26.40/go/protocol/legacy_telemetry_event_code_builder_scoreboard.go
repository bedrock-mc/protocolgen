// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventCodeBuilderScoreboard struct {
	ObjectiveName string
	Score         int32
}

func (LegacyTelemetryEventCodeBuilderScoreboard) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventCodeBuilderScoreboard using its canonical wire layout.
func (x *LegacyTelemetryEventCodeBuilderScoreboard) Marshal(io IO) {
	io.String(&x.ObjectiveName)
	io.Varint32(&x.Score)
}
