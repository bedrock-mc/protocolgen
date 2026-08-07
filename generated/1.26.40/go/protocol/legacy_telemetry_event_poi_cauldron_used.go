// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventPOICauldronUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemId               int32
}

func (LegacyTelemetryEventPOICauldronUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventPOICauldronUsed using its canonical wire layout.
func (x *LegacyTelemetryEventPOICauldronUsed) Marshal(io IO) {
	IntegerFunc(&x.BlockInteractionType, io.Uint8)
	io.Varint32(&x.ItemId)
}
