// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventCodeBuilderRuntimeAction struct {
	CodeBuilderRuntimeAction string
}

func (LegacyTelemetryEventCodeBuilderRuntimeAction) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventCodeBuilderRuntimeAction using its canonical wire layout.
func (x *LegacyTelemetryEventCodeBuilderRuntimeAction) Marshal(io IO) {
	io.String(&x.CodeBuilderRuntimeAction)
}
