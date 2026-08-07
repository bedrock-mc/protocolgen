// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MovementEffect struct {
	TargetRuntimeID uint64
	EffectID        MovementEffectType
	EffectDuration  int32
	Tick            uint64
}

// Marshal reads or writes MovementEffect using its canonical wire layout.
func (x *MovementEffect) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	IntegerFunc(&x.EffectID, io.Varint32)
	io.Varint32(&x.EffectDuration)
	io.PlayerInputTick(&x.Tick)
}
