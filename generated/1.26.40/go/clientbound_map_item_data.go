// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundMapItemData struct {
	MapID           int64
	Dimension       uint8
	IsLocked        bool
	MapOrigin       BlockPos
	CreationMapIDs  Optional[[]int64]
	Scale           Optional[int8]
	TrackedActorIDs Optional[[]MapItemTrackedActorUniqueId]
	Decorations     Optional[[]MapDecoration]
	Width           Optional[int32]
	Height          Optional[int32]
	StartX          Optional[int32]
	StartY          Optional[int32]
	Pixels          Optional[[]uint32]
}

// Marshal reads or writes ClientboundMapItemData using its canonical wire layout.
func (x *ClientboundMapItemData) Marshal(io IO) {
	io.ActorUniqueID(&x.MapID)
	io.Uint8(&x.Dimension)
	io.Bool(&x.IsLocked)
	x.MapOrigin.Marshal(io)
	OptionalFunc(io, &x.CreationMapIDs, func(value *[]int64) {
		FuncSlice(io, value, io.Varuint32, io.ActorUniqueID)
	})
	OptionalFunc(io, &x.Scale, io.Int8)
	OptionalFunc(io, &x.TrackedActorIDs, func(value *[]MapItemTrackedActorUniqueId) {
		FuncSlice(io, value, io.Varuint32, func(value *MapItemTrackedActorUniqueId) {
			value.Marshal(io)
		})
	})
	OptionalFunc(io, &x.Decorations, func(value *[]MapDecoration) {
		FuncSlice(io, value, io.Varuint32, func(value *MapDecoration) {
			value.Marshal(io)
		})
	})
	OptionalFunc(io, &x.Width, io.Varint32)
	OptionalFunc(io, &x.Height, io.Varint32)
	OptionalFunc(io, &x.StartX, io.Varint32)
	OptionalFunc(io, &x.StartY, io.Varint32)
	OptionalFunc(io, &x.Pixels, func(value *[]uint32) {
		FuncSlice(io, value, io.Varuint32, io.Uint32)
	})
}
