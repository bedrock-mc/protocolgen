// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// RequestAbility is a packet sent by the client to the server to request permission for a specific
// ability from the server. These abilities are defined above.
type RequestAbility struct {
	// Ability is the ability that the client is requesting. This is one of the constants defined in the
	// protocol/ability.go file.
	Ability   int32
	ValueType protocol.RequestAbilityType
	Bool      bool
	Float     float32
}

// Marshal reads or writes RequestAbility using its canonical wire layout.
func (x *RequestAbility) Marshal(io protocol.IO) {
	io.Varint32(&x.Ability)
	protocol.Minimum(io, &x.Ability, 0)
	protocol.Maximum(io, &x.Ability, 19)
	protocol.IntegerFunc(&x.ValueType, io.Uint8)
	io.Bool(&x.Bool)
	io.Float32(&x.Float)
}

// ID returns the protocol ID for RequestAbility.
func (*RequestAbility) ID() uint32 { return IDRequestAbility }
