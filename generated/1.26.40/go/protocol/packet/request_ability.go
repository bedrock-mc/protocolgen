// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RequestAbility struct {
	Ability   int32
	ValueType protocol.RequestAbilityType
	Bool      bool
	Float     float32
}

// Marshal reads or writes RequestAbility using its canonical wire layout.
func (x *RequestAbility) Marshal(io protocol.IO) {
	io.Varint32(&x.Ability)
	protocol.IntegerFunc(&x.ValueType, io.Uint8)
	io.Bool(&x.Bool)
	io.Float32(&x.Float)
}
