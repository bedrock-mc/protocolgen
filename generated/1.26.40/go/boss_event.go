// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BossEvent struct {
	TargetActorID ActorUniqueID
	PlayerID      ActorUniqueID
	EventType     BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         BossBarColor
	Overlay       BossBarOverlay
}

// Marshal reads or writes BossEvent using its canonical wire layout.
func (x *BossEvent) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	x.PlayerID.Marshal(io)
	enumValue1 := uint8(x.EventType)
	io.Uint8(&enumValue1)
	x.EventType = BossEventUpdateType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.Name)
	io.String(&x.FilteredName)
	io.Float32(&x.HealthPercent)
	enumValue2 := uint8(x.Color)
	io.Uint8(&enumValue2)
	x.Color = BossBarColor(enumValue2)
	switch int64(enumValue2) {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	enumValue3 := uint8(x.Overlay)
	io.Uint8(&enumValue3)
	x.Overlay = BossBarOverlay(enumValue3)
	switch int64(enumValue3) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue3, "unknown enum value")
	}
}
