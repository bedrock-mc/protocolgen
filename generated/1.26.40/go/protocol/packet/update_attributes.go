// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// UpdateAttributes is sent by the server to update an amount of attributes of any entity in the
// world. These attributes include ones such as the health or the movement speed of the entity.
type UpdateAttributes struct {
	TargetRuntimeID uint64
	AttributeList   []protocol.AttributeData
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// Marshal reads or writes UpdateAttributes using its canonical wire layout.
func (x *UpdateAttributes) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.Slice(io, &x.AttributeList)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for UpdateAttributes.
func (*UpdateAttributes) ID() uint32 { return IDUpdateAttributes }
