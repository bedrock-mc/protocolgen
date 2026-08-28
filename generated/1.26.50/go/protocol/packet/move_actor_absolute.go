// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// MoveActorAbsolute is sent by the server to move an entity to an absolute position. It is
// typically used for movements where high accuracy isn't needed, such as for long range
// teleporting.
type MoveActorAbsolute struct {
	MoveData protocol.MoveActorAbsoluteData
}

// Marshal reads or writes MoveActorAbsolute using its canonical wire layout.
func (x *MoveActorAbsolute) Marshal(io protocol.IO) {
	x.MoveData.Marshal(io)
}

// ID returns the protocol ID for MoveActorAbsolute.
func (*MoveActorAbsolute) ID() uint32 { return IDMoveActorAbsolute }
