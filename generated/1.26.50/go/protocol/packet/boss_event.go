// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// BossEvent is sent by the server to make a specific 'boss event' occur in the world. It includes
// features such as showing a boss bar to the player and turning the sky dark.
type BossEvent struct {
	TargetActorID int64
	PlayerID      int64
	// EventType is the type of the event. It is one of the BossEvent constants above.
	EventType     protocol.BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         protocol.BossBarColor
	// Overlay is the overlay of the boss bar that is shown on top of the boss bar when a player is
	// subscribed. It is one of the BossEventOverlay constants listed above.
	Overlay protocol.BossBarOverlay
}

// Marshal reads or writes BossEvent using its canonical wire layout.
func (x *BossEvent) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorUniqueID(&x.PlayerID)
	protocol.IntegerFunc(&x.EventType, io.Uint8)
	io.StringLimits(&x.Name, 0, 256)
	io.StringLimits(&x.FilteredName, 0, 256)
	io.Float32(&x.HealthPercent)
	protocol.IntegerFunc(&x.Color, io.Uint8)
	protocol.IntegerFunc(&x.Overlay, io.Uint8)
}

// ID returns the protocol ID for BossEvent.
func (*BossEvent) ID() uint32 { return IDBossEvent }
