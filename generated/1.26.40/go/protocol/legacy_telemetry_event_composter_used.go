// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventComposterUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemId               int32
}

func (LegacyTelemetryEventComposterUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventComposterUsed using its canonical wire layout.
func (x *LegacyTelemetryEventComposterUsed) Marshal(io IO) {
	IntegerFunc(&x.BlockInteractionType, io.Uint8)
	io.Varint32(&x.ItemId)
}
