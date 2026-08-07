// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type GraphicsOverrideParameter struct {
	ParameterKeyframeValues []OrderedEntry[float32, mgl32.Vec3]
	FloatValue              Optional[float32]
	Vec3Value               Optional[mgl32.Vec3]
	BiomeIdentifier         string
	PlayerIdentifier        Optional[string]
	IdentifierForParameter  GraphicsOverrideParameterType
	ResetParameter          bool
}

// Marshal reads or writes GraphicsOverrideParameter using its canonical wire layout.
func (x *GraphicsOverrideParameter) Marshal(io IO) {
	OrderedMap(io, &x.ParameterKeyframeValues, io.Varuint32, io.Float32, io.Vec3)
	OptionalFunc(io, &x.FloatValue, io.Float32)
	OptionalFunc(io, &x.Vec3Value, io.Vec3)
	io.String(&x.BiomeIdentifier)
	OptionalFunc(io, &x.PlayerIdentifier, io.String)
	IntegerFunc(&x.IdentifierForParameter, io.Uint8)
	io.Bool(&x.ResetParameter)
}
