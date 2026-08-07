// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper struct {
	PlayerWaxedOrUnwaxedCopperBlockID int32
}

func (LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper using its canonical wire layout.
func (x *LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper) Marshal(io IO) {
	io.Varint32(&x.PlayerWaxedOrUnwaxedCopperBlockID)
}
