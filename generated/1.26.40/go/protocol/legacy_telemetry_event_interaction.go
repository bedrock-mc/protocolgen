// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventInteraction struct {
	InteractedEntityID      int64
	InteractionType         MinecraftEventingInteractionType
	InteractionActorType    int32
	InteractionActorVariant int32
	InteractionActorColor   uint8
}

func (*LegacyTelemetryEventInteraction) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventInteraction using its canonical wire layout.
func (x *LegacyTelemetryEventInteraction) Marshal(io IO) {
	io.Varint64(&x.InteractedEntityID)
	IntegerFunc(&x.InteractionType, io.Uint8)
	io.Varint32(&x.InteractionActorType)
	io.Varint32(&x.InteractionActorVariant)
	io.Uint8(&x.InteractionActorColor)
}
