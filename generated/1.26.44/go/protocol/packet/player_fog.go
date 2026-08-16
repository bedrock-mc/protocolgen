// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// PlayerFog is sent by the server to render the different fogs in the Stack. The types of fog are
// controlled by resource packs to change how they are rendered, and the ability to create custom
// fog.
type PlayerFog struct {
	// FogStack is a list of fog identifiers to be sent to the client. Examples of fog identifiers are
	// "minecraft:fog_ocean" and "minecraft:fog_hell".
	FogStack []string
}

// Marshal reads or writes PlayerFog using its canonical wire layout.
func (x *PlayerFog) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.FogStack, io.Varuint32, io.String)
}

// ID returns the protocol ID for PlayerFog.
func (*PlayerFog) ID() uint32 { return IDPlayerFog }
