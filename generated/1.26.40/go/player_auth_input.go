// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type PlayerAuthInput struct {
	PlayerRotation         mgl32.Vec2
	Position               mgl32.Vec3
	MoveVector             mgl32.Vec2
	PlayerHeadRotation     float32
	InputData              Optional[[]PlayerAuthInputInputData]
	InputMode              InputMode
	PlayMode               ClientPlayMode
	NewInteractionModel    NewInteractionModel
	InteractRotation       mgl32.Vec2
	ClientTick             PlayerInputTick
	PosDelta               mgl32.Vec3
	ItemUseTransaction     Optional[PackedItemUseLegacyInventoryTransaction]
	ItemStackRequest       Optional[ItemStackRequestCerealRequestData]
	PlayerBlockActions     Optional[[]PlayerBlockActionData]
	VehicleRotation        Optional[mgl32.Vec2]
	ClientPredictedVehicle Optional[ActorUniqueID]
	AnalogMoveVector       mgl32.Vec2
	CameraOrientation      mgl32.Vec3
	RawMoveVector          mgl32.Vec2
}

// Marshal reads or writes PlayerAuthInput using its canonical wire layout.
func (x *PlayerAuthInput) Marshal(io IO) {
	io.Vec2(&x.PlayerRotation)
	io.Vec3(&x.Position)
	io.Vec2(&x.MoveVector)
	io.Float32(&x.PlayerHeadRotation)
	io.Bool(&x.InputData.set)
	if x.InputData.set {
		if !io.Reading() && uint64(len(x.InputData.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.InputData.val), "collection length overflows uint32")
			return
		}
		count1 := uint32(len(x.InputData.val))
		io.Varuint32(&count1)
		if io.Reading() {
			if uint64(count1) > uint64(^uint(0)>>1) {
				io.InvalidValue(count1, "collection length overflows int")
				return
			}
			x.InputData.val = make([]PlayerAuthInputInputData, int(count1))
		}
		for index2 := range x.InputData.val {
			enumValue3 := int32(x.InputData.val[index2])
			io.Varint32(&enumValue3)
			x.InputData.val[index2] = PlayerAuthInputInputData(enumValue3)
			switch int64(enumValue3) {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65:
			default:
				io.InvalidValue(enumValue3, "unknown enum value")
			}
		}
	} else if io.Reading() {
		var zero []PlayerAuthInputInputData
		x.InputData.val = zero
	}
	enumValue4 := uint32(x.InputMode)
	io.Varuint32(&enumValue4)
	x.InputMode = InputMode(enumValue4)
	switch int64(enumValue4) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue4, "unknown enum value")
	}
	enumValue5 := uint32(x.PlayMode)
	io.Varuint32(&enumValue5)
	x.PlayMode = ClientPlayMode(enumValue5)
	switch int64(enumValue5) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		io.InvalidValue(enumValue5, "unknown enum value")
	}
	enumValue6 := int32(x.NewInteractionModel)
	io.Varint32(&enumValue6)
	x.NewInteractionModel = NewInteractionModel(enumValue6)
	switch int64(enumValue6) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue6, "unknown enum value")
	}
	io.Vec2(&x.InteractRotation)
	x.ClientTick.Marshal(io)
	io.Vec3(&x.PosDelta)
	outer7 := true
	io.Bool(&outer7)
	if outer7 {
		io.Bool(&x.ItemUseTransaction.set)
		if x.ItemUseTransaction.set {
			x.ItemUseTransaction.val.Marshal(io)
		} else if io.Reading() {
			var zero PackedItemUseLegacyInventoryTransaction
			x.ItemUseTransaction.val = zero
		}
	} else {
		x.ItemUseTransaction = Optional[PackedItemUseLegacyInventoryTransaction]{}
	}
	outer8 := true
	io.Bool(&outer8)
	if outer8 {
		io.Bool(&x.ItemStackRequest.set)
		if x.ItemStackRequest.set {
			x.ItemStackRequest.val.Marshal(io)
		} else if io.Reading() {
			var zero ItemStackRequestCerealRequestData
			x.ItemStackRequest.val = zero
		}
	} else {
		x.ItemStackRequest = Optional[ItemStackRequestCerealRequestData]{}
	}
	outer9 := true
	io.Bool(&outer9)
	if outer9 {
		io.Bool(&x.PlayerBlockActions.set)
		if x.PlayerBlockActions.set {
			if !io.Reading() && uint64(len(x.PlayerBlockActions.val)) > uint64(^uint32(0)) {
				io.InvalidValue(len(x.PlayerBlockActions.val), "collection length overflows uint32")
				return
			}
			count10 := uint32(len(x.PlayerBlockActions.val))
			io.Varuint32(&count10)
			if io.Reading() {
				if uint64(count10) > uint64(^uint(0)>>1) {
					io.InvalidValue(count10, "collection length overflows int")
					return
				}
				x.PlayerBlockActions.val = make([]PlayerBlockActionData, int(count10))
			}
			for index11 := range x.PlayerBlockActions.val {
				x.PlayerBlockActions.val[index11].Marshal(io)
			}
		} else if io.Reading() {
			var zero []PlayerBlockActionData
			x.PlayerBlockActions.val = zero
		}
	} else {
		x.PlayerBlockActions = Optional[[]PlayerBlockActionData]{}
	}
	outer12 := true
	io.Bool(&outer12)
	if outer12 {
		io.Bool(&x.VehicleRotation.set)
		if x.VehicleRotation.set {
			io.Vec2(&x.VehicleRotation.val)
		} else if io.Reading() {
			var zero mgl32.Vec2
			x.VehicleRotation.val = zero
		}
	} else {
		x.VehicleRotation = Optional[mgl32.Vec2]{}
	}
	outer13 := true
	io.Bool(&outer13)
	if outer13 {
		io.Bool(&x.ClientPredictedVehicle.set)
		if x.ClientPredictedVehicle.set {
			x.ClientPredictedVehicle.val.Marshal(io)
		} else if io.Reading() {
			var zero ActorUniqueID
			x.ClientPredictedVehicle.val = zero
		}
	} else {
		x.ClientPredictedVehicle = Optional[ActorUniqueID]{}
	}
	io.Vec2(&x.AnalogMoveVector)
	io.Vec3(&x.CameraOrientation)
	io.Vec2(&x.RawMoveVector)
}
