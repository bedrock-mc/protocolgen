// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ChangeMobProperty struct {
	ActorId              int64
	PropertyName         string
	BoolComponentValue   bool
	StringComponentValue string
	IntComponentValue    int32
	FloatComponentValue  float32
}

// Marshal reads or writes ChangeMobProperty using its canonical wire layout.
func (x *ChangeMobProperty) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.ActorId)
	io.String(&x.PropertyName)
	io.Bool(&x.BoolComponentValue)
	io.String(&x.StringComponentValue)
	io.Varint32(&x.IntComponentValue)
	io.Float32(&x.FloatComponentValue)
}
