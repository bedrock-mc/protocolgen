// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// ChangeMobProperty is a packet sent from the server to the client to change one of the properties of a mob
// client-side.
type ChangeMobProperty struct {
	ActorID              int64
	PropertyName         string
	BoolComponentValue   bool
	StringComponentValue string
	IntComponentValue    int32
	FloatComponentValue  float32
}

// ID returns the protocol ID for ChangeMobProperty.
func (*ChangeMobProperty) ID() uint32 { return IDChangeMobProperty }

// Marshal reads or writes ChangeMobProperty using its canonical wire layout.
func (pk *ChangeMobProperty) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.ActorID)
	io.String(&pk.PropertyName)
	io.Bool(&pk.BoolComponentValue)
	io.String(&pk.StringComponentValue)
	io.Varint32(&pk.IntComponentValue)
	io.Float32(&pk.FloatComponentValue)
}
