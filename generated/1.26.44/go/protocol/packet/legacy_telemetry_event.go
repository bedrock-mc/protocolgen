// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

type LegacyTelemetryEvent struct {
	TargetActorID int64
	EventType     protocol.LegacyTelemetryType
	UsePlayerID   bool
	EventData     protocol.EventData
}

// ID returns the protocol ID for LegacyTelemetryEvent.
func (*LegacyTelemetryEvent) ID() uint32 { return IDLegacyTelemetryEvent }

// Marshal reads or writes LegacyTelemetryEvent using its canonical wire layout.
func (pk *LegacyTelemetryEvent) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.TargetActorID)
	pk.EventType.Marshal(io)
	io.Bool(&pk.UsePlayerID)
	protocol.MarshalEventData(io, &pk.EventData)
}
