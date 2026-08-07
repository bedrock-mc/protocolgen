// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type GraphicsOverrideParameter struct {
	ParameterKeyframeValues []protocol.OrderedEntry[float32, mgl32.Vec3]
	FloatValue              protocol.Optional[float32]
	Vec3Value               protocol.Optional[mgl32.Vec3]
	BiomeIdentifier         string
	PlayerIdentifier        protocol.Optional[string]
	IdentifierForParameter  protocol.GraphicsOverrideParameterType
	ResetParameter          bool
}

// Marshal reads or writes GraphicsOverrideParameter using its canonical wire layout.
func (x *GraphicsOverrideParameter) Marshal(io protocol.IO) {
	protocol.OrderedMap(io, &x.ParameterKeyframeValues, io.Varuint32, io.Float32, io.Vec3)
	protocol.OptionalFunc(io, &x.FloatValue, io.Float32)
	protocol.OptionalFunc(io, &x.Vec3Value, io.Vec3)
	io.String(&x.BiomeIdentifier)
	protocol.OptionalFunc(io, &x.PlayerIdentifier, io.String)
	protocol.IntegerFunc(&x.IdentifierForParameter, io.Uint8)
	io.Bool(&x.ResetParameter)
}
