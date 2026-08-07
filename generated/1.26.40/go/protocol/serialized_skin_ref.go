// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "image/color"

type SerializedSkinRef struct {
	ID                           string
	PlayFabID                    string
	ResourcePatch                string
	ImageData                    SkinImage
	AnimatedImageData            []AnimatedImageData
	CapeImageData                SkinImage
	GeometryData                 string
	GeometryDataMinEngineVersion string
	AnimationData                string
	CapeID                       string
	FullID                       string
	ArmSize                      PersonaArmSizeType
	SkinColor                    color.RGBA
	PersonaPieces                []SerializedPersonaPieceHandle
	PieceTintColors              []OrderedEntry[string, TintMapColor]
	IsPremium                    bool
	IsPersona                    bool
	IsPersonaCapeOnClassicSkin   bool
	IsPrimaryUser                bool
	OverridesPlayerAppearance    bool
	TrustedSkinFlag              string
	ProfileHash                  string
}

// Marshal reads or writes SerializedSkinRef using its canonical wire layout.
func (x *SerializedSkinRef) Marshal(io IO) {
	io.String(&x.ID)
	io.String(&x.PlayFabID)
	io.String(&x.ResourcePatch)
	x.ImageData.Marshal(io)
	FuncSlice(io, &x.AnimatedImageData, io.Varuint32, func(value *AnimatedImageData) {
		value.Marshal(io)
	})
	x.CapeImageData.Marshal(io)
	io.String(&x.GeometryData)
	io.String(&x.GeometryDataMinEngineVersion)
	io.String(&x.AnimationData)
	io.String(&x.CapeID)
	io.String(&x.FullID)
	IntegerFunc(&x.ArmSize, io.Uint8)
	io.RGBA(&x.SkinColor)
	FuncSlice(io, &x.PersonaPieces, io.Varuint32, func(value *SerializedPersonaPieceHandle) {
		value.Marshal(io)
	})
	OrderedMap(io, &x.PieceTintColors, io.Varuint32, io.String, func(value *TintMapColor) {
		value.Marshal(io)
	})
	io.Bool(&x.IsPremium)
	io.Bool(&x.IsPersona)
	io.Bool(&x.IsPersonaCapeOnClassicSkin)
	io.Bool(&x.IsPrimaryUser)
	io.Bool(&x.OverridesPlayerAppearance)
	io.String(&x.TrustedSkinFlag)
	io.String(&x.ProfileHash)
}
