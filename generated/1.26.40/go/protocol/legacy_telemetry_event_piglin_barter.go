// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventPiglinBarter struct {
	ItemId                      int32
	WasTargetingBarteringPlayer bool
}

func (LegacyTelemetryEventPiglinBarter) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPiglinBarter using its canonical wire layout.
func (x *LegacyTelemetryEventPiglinBarter) Marshal(io IO) {
	io.Varint32(&x.ItemId)
	io.Bool(&x.WasTargetingBarteringPlayer)
}
