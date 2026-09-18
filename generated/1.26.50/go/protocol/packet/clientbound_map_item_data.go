// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

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

// ID ...
func (*ClientboundMapItemData) ID() uint32 {
	return IDClientboundMapItemData
}

func (pk *ClientboundMapItemData) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.MapID)
	io.Uint8(&pk.Dimension)
	io.Bool(&pk.IsLocked)
	pk.MapOrigin.Marshal(io)
	protocol.OptionalFunc(io, &pk.CreationMapIDs, func(value *[]int64) {
		protocol.FuncSliceLimits(io, value, io.Varuint32, 0, 65535, io.ActorUniqueID)
	})
	protocol.OptionalFunc(io, &pk.Scale, io.Int8)
	protocol.OptionalFunc(io, &pk.TrackedActorIDs, func(value *[]protocol.MapItemTrackedActorUniqueID) {
		protocol.SliceLimits(io, value, 0, 65535)
	})
	protocol.OptionalFunc(io, &pk.Decorations, func(value *[]protocol.MapDecoration) {
		protocol.SliceLimits(io, value, 0, 65535)
	})
	protocol.OptionalFunc(io, &pk.Width, io.Varint32)
	protocol.OptionalFunc(io, &pk.Height, io.Varint32)
	protocol.OptionalFunc(io, &pk.StartX, io.Varint32)
	protocol.OptionalFunc(io, &pk.StartY, io.Varint32)
	protocol.OptionalFunc(io, &pk.Pixels, func(value *[]uint32) {
		protocol.FuncSliceLimits(io, value, io.Varuint32, 0, 16384, io.Uint32)
	})
}
