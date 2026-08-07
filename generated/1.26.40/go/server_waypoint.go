// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
)

type ServerWaypoint struct {
	UpdateFlag              uint32
	IsVisible               Optional[bool]
	WorldPosition           Optional[WorldPosition]
	TexturePath             Optional[string]
	IconSize                Optional[mgl32.Vec2]
	Color                   Optional[color.RGBA]
	ClientPositionAuthority Optional[bool]
	ActorUniqueID           Optional[int64]
}

// Marshal reads or writes ServerWaypoint using its canonical wire layout.
func (x *ServerWaypoint) Marshal(io IO) {
	io.Uint32(&x.UpdateFlag)
	OptionalFunc(io, &x.IsVisible, io.Bool)
	OptionalFunc(io, &x.WorldPosition, func(value *WorldPosition) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.TexturePath, io.String)
	OptionalFunc(io, &x.IconSize, io.Vec2)
	OptionalFunc(io, &x.Color, io.RGBA)
	OptionalFunc(io, &x.ClientPositionAuthority, io.Bool)
	OptionalFunc(io, &x.ActorUniqueID, io.ActorUniqueID)
}
