// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type GraphicsOverrideParameter struct {
	ParameterKeyframeValues []OrderedEntry[float32, mgl32.Vec3]
	FloatValue              *float32
	Vec3Value               *mgl32.Vec3
	BiomeIdentifier         string
	PlayerIdentifier        *string
	IdentifierForParameter  GraphicsOverrideParameterType
	ResetParameter          bool
}
