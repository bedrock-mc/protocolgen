// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MovementEffect struct {
	TargetRuntimeID uint64
	EffectID        protocol.MovementEffectType
	EffectDuration  int32
	Tick            uint64
}

// Marshal reads or writes MovementEffect using its canonical wire layout.
func (x *MovementEffect) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.IntegerFunc(&x.EffectID, io.Varint32)
	io.Varint32(&x.EffectDuration)
	io.PlayerInputTick(&x.Tick)
}
