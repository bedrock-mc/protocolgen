// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// MoveActorDelta is sent by the server to move an entity. The packet is specifically optimised to
// save as much space as possible, by only writing non-zero fields. As of 1.16.100, this packet no
// longer actually contains any deltas.
type MoveActorDelta struct {
	MoveData protocol.MoveActorDeltaData
}

// Marshal reads or writes MoveActorDelta using its canonical wire layout.
func (x *MoveActorDelta) Marshal(io protocol.IO) {
	x.MoveData.Marshal(io)
}

// ID returns the protocol ID for MoveActorDelta.
func (*MoveActorDelta) ID() uint32 { return IDMoveActorDelta }
