// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerLocation struct {
	TargetActorID int64
	Location      protocol.PlayerLocationLocation
}

// Marshal reads or writes PlayerLocation using its canonical wire layout.
func (x *PlayerLocation) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	protocol.MarshalPlayerLocationLocation(io, &x.Location)
}

// ID returns the protocol ID for PlayerLocation.
func (*PlayerLocation) ID() uint32 { return IDPlayerLocation }
