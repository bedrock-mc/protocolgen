// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// UpdateAttributes is sent by the server to update an amount of attributes of any entity in the world. These
// attributes include ones such as the health or the movement speed of the entity.
type UpdateAttributes struct {
	TargetRuntimeID uint64
	AttributeList   []protocol.AttributeData
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// ID returns the protocol ID for UpdateAttributes.
func (*UpdateAttributes) ID() uint32 { return IDUpdateAttributes }

// Marshal reads or writes UpdateAttributes using its canonical wire layout.
func (pk *UpdateAttributes) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	protocol.Slice(io, &pk.AttributeList)
	io.PlayerInputTick(&pk.Tick)
}
