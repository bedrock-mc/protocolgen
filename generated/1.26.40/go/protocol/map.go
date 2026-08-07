// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "image/color"

// MapDecoration is a fixed decoration on a map: Its position or other properties do not change
// automatically client-side.
type MapDecoration struct {
	ImageType MapDecorationType
	// Rotation is the rotation of the map decoration. It is byte due to the 16 fixed directions that
	// the map decoration may face.
	Rotation uint8
	// X is the offset on the X axis in pixels of the decoration.
	X uint8
	// Y is the offset on the Y axis in pixels of the decoration.
	Y uint8
	// Label is the name of the map decoration. This name may be of any value.
	Label string
	Color color.RGBA
}

// Marshal reads or writes MapDecoration using its canonical wire layout.
func (x *MapDecoration) Marshal(io IO) {
	IntegerFunc(&x.ImageType, io.Int8)
	io.Uint8(&x.Rotation)
	io.Uint8(&x.X)
	io.Uint8(&x.Y)
	io.String(&x.Label)
	io.RGBA(&x.Color)
}

// MapDecoration is a fixed decoration on a map: Its position or other properties do not change
// automatically client-side.
type MapDecorationType int8

const (
	MapDecorationTypeMarkerWhite      MapDecorationType = 0
	MapDecorationTypeMarkerGreen      MapDecorationType = 1
	MapDecorationTypeMarkerRed        MapDecorationType = 2
	MapDecorationTypeMarkerBlue       MapDecorationType = 3
	MapDecorationTypeXWhite           MapDecorationType = 4
	MapDecorationTypeTriangleRed      MapDecorationType = 5
	MapDecorationTypeSquareWhite      MapDecorationType = 6
	MapDecorationTypeMarkerSign       MapDecorationType = 7
	MapDecorationTypeMarkerPink       MapDecorationType = 8
	MapDecorationTypeMarkerOrange     MapDecorationType = 9
	MapDecorationTypeMarkerYellow     MapDecorationType = 10
	MapDecorationTypeMarkerTeal       MapDecorationType = 11
	MapDecorationTypeTriangleGreen    MapDecorationType = 12
	MapDecorationTypeSmallSquareWhite MapDecorationType = 13
	MapDecorationTypeMansion          MapDecorationType = 14
	MapDecorationTypeMonument         MapDecorationType = 15
	MapDecorationTypeNoDraw           MapDecorationType = 16
	MapDecorationTypeVillageDesert    MapDecorationType = 17
	MapDecorationTypeVillagePlains    MapDecorationType = 18
	MapDecorationTypeVillageSavanna   MapDecorationType = 19
	MapDecorationTypeVillageSnowy     MapDecorationType = 20
	MapDecorationTypeVillageTaiga     MapDecorationType = 21
	MapDecorationTypeJungleTemple     MapDecorationType = 22
	MapDecorationTypeWitchHut         MapDecorationType = 23
	MapDecorationTypeTrialChambers    MapDecorationType = 24
	MapDecorationTypeCount            MapDecorationType = 25
)

type MapItemTrackedActorType int32

const (
	MapItemTrackedActorTypeEntity      MapItemTrackedActorType = 0
	MapItemTrackedActorTypeBlockEntity MapItemTrackedActorType = 1
	MapItemTrackedActorTypeOther       MapItemTrackedActorType = 2
)

type MapItemTrackedActorUniqueID struct {
	Type          MapItemTrackedActorType
	EntityID      Optional[int64]
	BlockPosition Optional[BlockPos]
}

// Marshal reads or writes MapItemTrackedActorUniqueID using its canonical wire layout.
func (x *MapItemTrackedActorUniqueID) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Int32)
	OptionalFunc(io, &x.EntityID, io.ActorUniqueID)
	OptionalFunc(io, &x.BlockPosition, func(value *BlockPos) {
		value.Marshal(io)
	})
}

// PixelRequest is the request for the colour of a pixel in a MapInfoRequest packet.
type PixelRequest struct {
	Pixel uint32
	Index uint16
}

// Marshal reads or writes PixelRequest using its canonical wire layout.
func (x *PixelRequest) Marshal(io IO) {
	io.Uint32(&x.Pixel)
	io.Uint16(&x.Index)
}
