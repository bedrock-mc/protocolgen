// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type ServerPlayerPostMovePosition struct {
	Pos mgl32.Vec3
}

// ID ...
func (*ServerPlayerPostMovePosition) ID() uint32 {
	return IDServerPlayerPostMovePosition
}

func (pk *ServerPlayerPostMovePosition) Marshal(io protocol.IO) {
	io.Vec3(&pk.Pos)
}
