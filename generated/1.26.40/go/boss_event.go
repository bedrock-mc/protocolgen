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
