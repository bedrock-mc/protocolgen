// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// PersonaPiece represents a piece of a persona skin. All pieces are sent separately.
type PersonaPieceType uint32

const (
	PersonaPieceTypeSkeleton      PersonaPieceType = 1
	PersonaPieceTypeBody          PersonaPieceType = 2
	PersonaPieceTypeSkin          PersonaPieceType = 3
	PersonaPieceTypeBottom        PersonaPieceType = 4
	PersonaPieceTypeFeet          PersonaPieceType = 5
	PersonaPieceTypeDress         PersonaPieceType = 6
	PersonaPieceTypeTop           PersonaPieceType = 7
	PersonaPieceTypeHighPants     PersonaPieceType = 8
	PersonaPieceTypeHands         PersonaPieceType = 9
	PersonaPieceTypeOuterwear     PersonaPieceType = 10
	PersonaPieceTypeFacialHair    PersonaPieceType = 11
	PersonaPieceTypeMouth         PersonaPieceType = 12
	PersonaPieceTypeEyes          PersonaPieceType = 13
	PersonaPieceTypeHair          PersonaPieceType = 14
	PersonaPieceTypeHood          PersonaPieceType = 15
	PersonaPieceTypeBack          PersonaPieceType = 16
	PersonaPieceTypeFaceAccessory PersonaPieceType = 17
	PersonaPieceTypeHead          PersonaPieceType = 18
	PersonaPieceTypeLegs          PersonaPieceType = 19
	PersonaPieceTypeLeftLeg       PersonaPieceType = 20
	PersonaPieceTypeRightLeg      PersonaPieceType = 21
	PersonaPieceTypeArms          PersonaPieceType = 22
	PersonaPieceTypeLeftArm       PersonaPieceType = 23
	PersonaPieceTypeRightArm      PersonaPieceType = 24
	PersonaPieceTypeCapes         PersonaPieceType = 25
	PersonaPieceTypeClassicSkin   PersonaPieceType = 26
	PersonaPieceTypeEmote         PersonaPieceType = 27
)

type SkinImage struct {
	Width      uint32
	Height     uint32
	ImageBytes []uint8
}

// Marshal reads or writes SkinImage using its canonical wire layout.
func (x *SkinImage) Marshal(io IO) {
	io.Uint32(&x.Width)
	Maximum(io, &x.Width, 4096)
	io.Uint32(&x.Height)
	Maximum(io, &x.Height, 4096)
	FuncSliceLimits(io, &x.ImageBytes, io.Varuint32, 0, 67108864, io.Uint8)
}
