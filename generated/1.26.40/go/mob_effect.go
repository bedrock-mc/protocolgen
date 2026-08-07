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
