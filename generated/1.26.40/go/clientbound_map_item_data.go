// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundMapItemData struct {
	MapID           ActorUniqueID
	Dimension       uint8
	IsLocked        bool
	MapOrigin       BlockPos
	CreationMapIDs  *[]ActorUniqueID
	Scale           *int8
	TrackedActorIDs *[]MapItemTrackedActorUniqueId
	Decorations     *[]MapDecoration
	Width           *int32
	Height          *int32
	StartX          *int32
	StartY          *int32
	Pixels          *[]uint32
}
