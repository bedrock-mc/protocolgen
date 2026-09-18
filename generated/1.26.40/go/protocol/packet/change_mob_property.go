// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
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

// ID ...
func (*ChangeMobProperty) ID() uint32 {
	return IDChangeMobProperty
}

func (pk *ChangeMobProperty) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.ActorID)
	io.String(&pk.PropertyName)
	io.Bool(&pk.BoolComponentValue)
	io.String(&pk.StringComponentValue)
	io.Varint32(&pk.IntComponentValue)
	io.Float32(&pk.FloatComponentValue)
}
