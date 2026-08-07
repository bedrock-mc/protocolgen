// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventItemUsed struct {
	ItemId    int16
	ItemAux   int32
	UseMethod int32
	Count     int32
}

func (LegacyTelemetryEventItemUsed) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventItemUsed using its canonical wire layout.
func (x *LegacyTelemetryEventItemUsed) Marshal(io IO) {
	io.Int16(&x.ItemId)
	io.Int32(&x.ItemAux)
	io.Int32(&x.UseMethod)
	io.Int32(&x.Count)
}
