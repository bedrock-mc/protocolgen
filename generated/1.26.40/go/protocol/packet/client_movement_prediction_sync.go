// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ClientMovementPredictionSync is sent by the client to the server periodically if the client has
// received movement corrections from the server, containing information about client-predictions
// that are relevant to movement.
type ClientMovementPredictionSync struct {
	ActorDataFlag      protocol.ActorDataFlagComponent
	ActorBoundingBox   protocol.ActorDataBoundingBoxComponent
	MovementAttributes [9]float32
	ActorUniqueID      int64
	ActorFlyingState   bool
}

// Marshal reads or writes ClientMovementPredictionSync using its canonical wire layout.
func (x *ClientMovementPredictionSync) Marshal(io protocol.IO) {
	x.ActorDataFlag.Marshal(io)
	x.ActorBoundingBox.Marshal(io)
	for index1 := range x.MovementAttributes {
		io.Float32(&x.MovementAttributes[index1])
	}
	io.ActorUniqueID(&x.ActorUniqueID)
	io.Bool(&x.ActorFlyingState)
}

// ID returns the protocol ID for ClientMovementPredictionSync.
func (*ClientMovementPredictionSync) ID() uint32 { return IDClientMovementPredictionSync }
