// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// MobEffect is sent by the server to apply an effect to the player, for example an effect like
// poison. It may also be used to modify existing effects, or removing them completely.
type MobEffect struct {
	TargetRuntimeID     uint64
	EventID             protocol.MobEffectEvent
	EffectID            int32
	EffectAmplifier     int32
	ShowParticles       bool
	EffectDurationTicks int32
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
	// Ambient specifies if the effect is ambient. If set to false, it will not get treated as an
	// ambient effect.
	Ambient bool
}

// Marshal reads or writes MobEffect using its canonical wire layout.
func (x *MobEffect) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.IntegerFunc(&x.EventID, io.Uint8)
	io.Varint32(&x.EffectID)
	io.Varint32(&x.EffectAmplifier)
	io.Bool(&x.ShowParticles)
	io.Varint32(&x.EffectDurationTicks)
	io.PlayerInputTick(&x.Tick)
	io.Bool(&x.Ambient)
}

// ID returns the protocol ID for MobEffect.
func (*MobEffect) ID() uint32 { return IDMobEffect }
