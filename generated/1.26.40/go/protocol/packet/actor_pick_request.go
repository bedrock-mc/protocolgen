// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ActorPickRequest struct {
	ActorID  int64
	MaxSlots uint8
	WithData bool
}

// Marshal reads or writes ActorPickRequest using its canonical wire layout.
func (x *ActorPickRequest) Marshal(io protocol.IO) {
	io.Int64(&x.ActorID)
	io.Uint8(&x.MaxSlots)
	io.Bool(&x.WithData)
}

// ID returns the protocol ID for ActorPickRequest.
func (*ActorPickRequest) ID() uint32 { return IDActorPickRequest }
