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
	if !io.Reading() && uint64(len(x.ParameterKeyframeValues)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ParameterKeyframeValues), "map length overflows uint32")
		return
	}
	count1 := uint32(len(x.ParameterKeyframeValues))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "map length overflows int")
			return
		}
		x.ParameterKeyframeValues = make([]OrderedEntry[float32, mgl32.Vec3], int(count1))
	}
	for index2 := range x.ParameterKeyframeValues {
		io.Float32(&x.ParameterKeyframeValues[index2].Key)
		io.Vec3(&x.ParameterKeyframeValues[index2].Value)
	}
	io.Bool(&x.FloatValue.set)
	if x.FloatValue.set {
		io.Float32(&x.FloatValue.val)
	} else if io.Reading() {
		var zero float32
		x.FloatValue.val = zero
	}
	io.Bool(&x.Vec3Value.set)
	if x.Vec3Value.set {
		io.Vec3(&x.Vec3Value.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.Vec3Value.val = zero
	}
	io.String(&x.BiomeIdentifier)
	io.Bool(&x.PlayerIdentifier.set)
	if x.PlayerIdentifier.set {
		io.String(&x.PlayerIdentifier.val)
	} else if io.Reading() {
		var zero string
		x.PlayerIdentifier.val = zero
	}
	enumValue3 := uint8(x.IdentifierForParameter)
	io.Uint8(&enumValue3)
	x.IdentifierForParameter = GraphicsOverrideParameterType(enumValue3)
	switch int64(enumValue3) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51:
	default:
		io.InvalidValue(enumValue3, "unknown enum value")
	}
	io.Bool(&x.ResetParameter)
}
