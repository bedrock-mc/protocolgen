// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LegacyTelemetryEvent struct {
	TargetActorID int64
	EventType     protocol.LegacyTelemetryType
	UsePlayerID   bool
	EventData     protocol.EventData
}

// Marshal reads or writes LegacyTelemetryEvent using its canonical wire layout.
func (x *LegacyTelemetryEvent) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	protocol.IntegerFunc(&x.EventType, io.Varint32)
	io.Bool(&x.UsePlayerID)
	protocol.MarshalEventData(io, &x.EventData)
}

// ID returns the protocol ID for LegacyTelemetryEvent.
func (*LegacyTelemetryEvent) ID() uint32 { return IDLegacyTelemetryEvent }
