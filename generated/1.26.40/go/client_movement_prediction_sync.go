// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientMovementPredictionSync struct {
	ActorDataFlag      ActorDataFlagComponent
	ActorBoundingBox   ActorDataBoundingBoxComponent
	MovementAttributes [9]float32
	ActorUniqueID      int64
	ActorFlyingState   bool
}

// Marshal reads or writes ClientMovementPredictionSync using its canonical wire layout.
func (x *ClientMovementPredictionSync) Marshal(io IO) {
	x.ActorDataFlag.Marshal(io)
	x.ActorBoundingBox.Marshal(io)
	for index1 := range x.MovementAttributes {
		io.Float32(&x.MovementAttributes[index1])
	}
	io.ActorUniqueID(&x.ActorUniqueID)
	io.Bool(&x.ActorFlyingState)
}
