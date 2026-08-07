// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AnimatedImageData struct {
	SkinImage           SkinImage
	AnimatedTextureType PersonaAnimatedTextureType
	Frames              float32
	AnimationExpression PersonaAnimationExpression
}

// Marshal reads or writes AnimatedImageData using its canonical wire layout.
func (x *AnimatedImageData) Marshal(io IO) {
	x.SkinImage.Marshal(io)
	IntegerFunc(&x.AnimatedTextureType, io.Varuint32)
	io.Float32(&x.Frames)
	IntegerFunc(&x.AnimationExpression, io.Varuint32)
}
