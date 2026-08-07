// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// GraphicsOverrideParameter is sent by the server to override graphics parameters.
type GraphicsOverrideParameter struct {
	ParameterKeyframeValues []protocol.OrderedEntry[float32, mgl32.Vec3]
	// FloatValue is an optional single float graphics parameter to be overridden.
	FloatValue protocol.Optional[float32]
	// Vec3Value is an optional single Vec3 graphics parameter to be overridden.
	Vec3Value protocol.Optional[mgl32.Vec3]
	// BiomeIdentifier is the identifier of the biome for which the parameters apply.
	BiomeIdentifier string
	// PlayerIdentifier is the optional identifier of the player for which the override parameter
	// applies.
	PlayerIdentifier       protocol.Optional[string]
	IdentifierForParameter protocol.GraphicsOverrideParameterType
	ResetParameter         bool
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

// ID returns the protocol ID for GraphicsOverrideParameter.
func (*GraphicsOverrideParameter) ID() uint32 { return IDGraphicsOverrideParameter }
