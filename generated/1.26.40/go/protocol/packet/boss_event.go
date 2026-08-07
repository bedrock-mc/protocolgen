// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type BossEvent struct {
	TargetActorID int64
	PlayerID      int64
	EventType     protocol.BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         protocol.BossBarColor
	Overlay       protocol.BossBarOverlay
}

// Marshal reads or writes BossEvent using its canonical wire layout.
func (x *BossEvent) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorUniqueID(&x.PlayerID)
	protocol.IntegerFunc(&x.EventType, io.Uint8)
	io.String(&x.Name)
	io.String(&x.FilteredName)
	io.Float32(&x.HealthPercent)
	protocol.IntegerFunc(&x.Color, io.Uint8)
	protocol.IntegerFunc(&x.Overlay, io.Uint8)
}
