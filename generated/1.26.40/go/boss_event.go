// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BossEvent struct {
	TargetActorID int64
	PlayerID      int64
	EventType     BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         BossBarColor
	Overlay       BossBarOverlay
}

// Marshal reads or writes BossEvent using its canonical wire layout.
func (x *BossEvent) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorUniqueID(&x.PlayerID)
	IntegerFunc(&x.EventType, io.Uint8)
	io.String(&x.Name)
	io.String(&x.FilteredName)
	io.Float32(&x.HealthPercent)
	IntegerFunc(&x.Color, io.Uint8)
	IntegerFunc(&x.Overlay, io.Uint8)
}
