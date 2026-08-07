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
