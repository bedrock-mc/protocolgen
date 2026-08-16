// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ActorPickRequest is sent by the client when it tries to pick an entity, so that it gets a spawn
// egg which can spawn that entity.
type ActorPickRequest struct {
	ActorID  int64
	MaxSlots uint8
	// WithData is true if the pick request requests the entity metadata.
	WithData bool
}

// Marshal reads or writes ActorPickRequest using its canonical wire layout.
func (x *ActorPickRequest) Marshal(io protocol.IO) {
	io.Int64(&x.ActorID)
	io.Uint8(&x.MaxSlots)
	protocol.Minimum(io, &x.MaxSlots, 0)
	protocol.Maximum(io, &x.MaxSlots, 255)
	io.Bool(&x.WithData)
}

// ID returns the protocol ID for ActorPickRequest.
func (*ActorPickRequest) ID() uint32 { return IDActorPickRequest }
