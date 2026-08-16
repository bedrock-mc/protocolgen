// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type ServerPlayerPostMovePosition struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes ServerPlayerPostMovePosition using its canonical wire layout.
func (x *ServerPlayerPostMovePosition) Marshal(io protocol.IO) {
	io.Vec3(&x.Pos)
}

// ID returns the protocol ID for ServerPlayerPostMovePosition.
func (*ServerPlayerPostMovePosition) ID() uint32 { return IDServerPlayerPostMovePosition }
