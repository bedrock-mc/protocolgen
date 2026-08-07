// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventBellUsed struct {
	ItemId int32
}

func (*LegacyTelemetryEventBellUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventBellUsed using its canonical wire layout.
func (x *LegacyTelemetryEventBellUsed) Marshal(io IO) {
	io.Varint32(&x.ItemId)
}
