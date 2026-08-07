// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventSlashCommand struct {
	SuccessCount int32
	ErrorCount   int32
	CommandName  string
	ErrorList    string
}

func (*LegacyTelemetryEventSlashCommand) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventSlashCommand using its canonical wire layout.
func (x *LegacyTelemetryEventSlashCommand) Marshal(io IO) {
	io.Varint32(&x.SuccessCount)
	io.Varint32(&x.ErrorCount)
	io.String(&x.CommandName)
	io.String(&x.ErrorList)
}
