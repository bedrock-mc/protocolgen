// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MobEffect struct {
	TargetRuntimeID     uint64
	EventID             MobEffectEvent
	EffectID            int32
	EffectAmplifier     int32
	ShowParticles       bool
	EffectDurationTicks int32
	Tick                uint64
	Ambient             bool
}

// Marshal reads or writes MobEffect using its canonical wire layout.
func (x *MobEffect) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	IntegerFunc(&x.EventID, io.Uint8)
	io.Varint32(&x.EffectID)
	io.Varint32(&x.EffectAmplifier)
	io.Bool(&x.ShowParticles)
	io.Varint32(&x.EffectDurationTicks)
	io.PlayerInputTick(&x.Tick)
	io.Bool(&x.Ambient)
}
