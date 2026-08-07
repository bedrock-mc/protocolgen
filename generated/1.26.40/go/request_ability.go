// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RequestAbility struct {
	Ability   int32
	ValueType RequestAbilityType
	Bool      bool
	Float     float32
}

// Marshal reads or writes RequestAbility using its canonical wire layout.
func (x *RequestAbility) Marshal(io IO) {
	io.Varint32(&x.Ability)
	IntegerFunc(&x.ValueType, io.Uint8)
	io.Bool(&x.Bool)
	io.Float32(&x.Float)
}
