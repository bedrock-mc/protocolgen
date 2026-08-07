// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundMapItemData struct {
	MapID           int64
	Dimension       uint8
	IsLocked        bool
	MapOrigin       protocol.BlockPos
	CreationMapIDs  protocol.Optional[[]int64]
	Scale           protocol.Optional[int8]
	TrackedActorIDs protocol.Optional[[]protocol.MapItemTrackedActorUniqueID]
	Decorations     protocol.Optional[[]protocol.MapDecoration]
	Width           protocol.Optional[int32]
	Height          protocol.Optional[int32]
	StartX          protocol.Optional[int32]
	StartY          protocol.Optional[int32]
	Pixels          protocol.Optional[[]uint32]
}

// Marshal reads or writes ClientboundMapItemData using its canonical wire layout.
func (x *ClientboundMapItemData) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.MapID)
	io.Uint8(&x.Dimension)
	io.Bool(&x.IsLocked)
	x.MapOrigin.Marshal(io)
	protocol.OptionalFunc(io, &x.CreationMapIDs, func(value *[]int64) {
		protocol.FuncSlice(io, value, io.Varuint32, io.ActorUniqueID)
	})
	protocol.OptionalFunc(io, &x.Scale, io.Int8)
	protocol.OptionalFunc(io, &x.TrackedActorIDs, func(value *[]protocol.MapItemTrackedActorUniqueID) {
		protocol.Slice(io, value)
	})
	protocol.OptionalFunc(io, &x.Decorations, func(value *[]protocol.MapDecoration) {
		protocol.Slice(io, value)
	})
	protocol.OptionalFunc(io, &x.Width, io.Varint32)
	protocol.OptionalFunc(io, &x.Height, io.Varint32)
	protocol.OptionalFunc(io, &x.StartX, io.Varint32)
	protocol.OptionalFunc(io, &x.StartY, io.Varint32)
	protocol.OptionalFunc(io, &x.Pixels, func(value *[]uint32) {
		protocol.FuncSlice(io, value, io.Varuint32, io.Uint32)
	})
}

// ID returns the protocol ID for ClientboundMapItemData.
func (*ClientboundMapItemData) ID() uint32 { return IDClientboundMapItemData }
