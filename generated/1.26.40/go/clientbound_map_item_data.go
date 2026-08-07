// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundMapItemData struct {
	MapID           ActorUniqueID
	Dimension       uint8
	IsLocked        bool
	MapOrigin       BlockPos
	CreationMapIDs  Optional[[]ActorUniqueID]
	Scale           Optional[int8]
	TrackedActorIDs Optional[[]MapItemTrackedActorUniqueId]
	Decorations     Optional[[]MapDecoration]
	Width           Optional[int32]
	Height          Optional[int32]
	StartX          Optional[int32]
	StartY          Optional[int32]
	Pixels          Optional[[]uint32]
}
