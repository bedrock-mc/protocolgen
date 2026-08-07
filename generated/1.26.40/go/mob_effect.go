// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MobEffect struct {
	TargetRuntimeID     ActorRuntimeID
	EventID             MobEffectEvent
	EffectID            int32
	EffectAmplifier     int32
	ShowParticles       bool
	EffectDurationTicks int32
	Tick                PlayerInputTick
	Ambient             bool
}

// Marshal reads or writes MobEffect using its canonical wire layout.
func (x *MobEffect) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	enumValue1 := uint8(x.EventID)
	io.Uint8(&enumValue1)
	x.EventID = MobEffectEvent(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Varint32(&x.EffectID)
	io.Varint32(&x.EffectAmplifier)
	io.Bool(&x.ShowParticles)
	io.Varint32(&x.EffectDurationTicks)
	x.Tick.Marshal(io)
	io.Bool(&x.Ambient)
}
