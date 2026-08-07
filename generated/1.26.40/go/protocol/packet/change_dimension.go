// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type ChangeDimension struct {
	DimensionID     protocol.DimensionType
	Position        mgl32.Vec3
	Respawn         bool
	LoadingScreenID protocol.Optional[uint32]
}

// Marshal reads or writes ChangeDimension using its canonical wire layout.
func (x *ChangeDimension) Marshal(io protocol.IO) {
	x.DimensionID.Marshal(io)
	io.Vec3(&x.Position)
	io.Bool(&x.Respawn)
	protocol.OptionalFunc(io, &x.LoadingScreenID, io.Uint32)
}

// ID returns the protocol ID for ChangeDimension.
func (*ChangeDimension) ID() uint32 { return IDChangeDimension }
