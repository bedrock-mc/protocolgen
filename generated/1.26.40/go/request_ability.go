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
	enumValue1 := uint8(x.ValueType)
	io.Uint8(&enumValue1)
	x.ValueType = RequestAbilityType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.Bool)
	io.Float32(&x.Float)
}
