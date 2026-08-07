// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventPortalUsed struct {
	SourceDimensionID int32
	TargetDimensionID int32
}

func (*LegacyTelemetryEventPortalUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPortalUsed using its canonical wire layout.
func (x *LegacyTelemetryEventPortalUsed) Marshal(io IO) {
	io.Varint32(&x.SourceDimensionID)
	io.Varint32(&x.TargetDimensionID)
}
