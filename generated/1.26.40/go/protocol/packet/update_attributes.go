// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateAttributes struct {
	TargetRuntimeID uint64
	AttributeList   []protocol.AttributeData
	Tick            uint64
}

// Marshal reads or writes UpdateAttributes using its canonical wire layout.
func (x *UpdateAttributes) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.Slice(io, &x.AttributeList)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for UpdateAttributes.
func (*UpdateAttributes) ID() uint32 { return IDUpdateAttributes }
