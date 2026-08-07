// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventCauldronUsed struct {
	ContentsColor uint32
	ContentsType  int32
	FillLevel     int32
}

func (LegacyTelemetryEventCauldronUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventCauldronUsed using its canonical wire layout.
func (x *LegacyTelemetryEventCauldronUsed) Marshal(io IO) {
	io.Varuint32(&x.ContentsColor)
	io.Varint32(&x.ContentsType)
	io.Varint32(&x.FillLevel)
}
