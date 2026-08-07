// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MovementEffect struct {
	TargetRuntimeID ActorRuntimeID
	EffectID        MovementEffectType
	EffectDuration  int32
	Tick            PlayerInputTick
}
