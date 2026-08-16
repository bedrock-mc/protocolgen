// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// PlayerLocation is sent by the server to the client to either update a player's position on the
// locator bar, or remove them completely. The client will determine how to render the player on the
// locator bar based on their own distance to Position.
type PlayerLocation struct {
	TargetActorID int64
	Location      protocol.PlayerLocationData
}

// Marshal reads or writes PlayerLocation using its canonical wire layout.
func (x *PlayerLocation) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	protocol.MarshalPlayerLocationData(io, &x.Location)
}

// ID returns the protocol ID for PlayerLocation.
func (*PlayerLocation) ID() uint32 { return IDPlayerLocation }
