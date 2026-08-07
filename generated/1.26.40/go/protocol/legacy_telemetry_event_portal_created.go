// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventPortalCreated struct {
	DimensionID int32
}

func (LegacyTelemetryEventPortalCreated) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPortalCreated using its canonical wire layout.
func (x *LegacyTelemetryEventPortalCreated) Marshal(io IO) {
	io.Varint32(&x.DimensionID)
}
