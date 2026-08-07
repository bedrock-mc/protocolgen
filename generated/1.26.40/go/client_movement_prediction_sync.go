// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientMovementPredictionSync struct {
	ActorDataFlag      ActorDataFlagComponent
	ActorBoundingBox   ActorDataBoundingBoxComponent
	MovementAttributes [9]float32
	ActorUniqueID      ActorUniqueID
	ActorFlyingState   bool
}
