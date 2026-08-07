// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventTargetBlockHit struct {
	RedstoneLevel int32
}

func (LegacyTelemetryEventTargetBlockHit) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventTargetBlockHit using its canonical wire layout.
func (x *LegacyTelemetryEventTargetBlockHit) Marshal(io IO) {
	io.Varint32(&x.RedstoneLevel)
}
