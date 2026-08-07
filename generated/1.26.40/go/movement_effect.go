// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MovementEffect struct {
	TargetRuntimeID ActorRuntimeID
	EffectID        MovementEffectType
	EffectDuration  int32
	Tick            PlayerInputTick
}

// Marshal reads or writes MovementEffect using its canonical wire layout.
func (x *MovementEffect) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	enumValue1 := int32(x.EffectID)
	io.Varint32(&enumValue1)
	x.EffectID = MovementEffectType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Varint32(&x.EffectDuration)
	x.Tick.Marshal(io)
}
