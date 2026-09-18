// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// MobEffect is sent by the server to apply an effect to the player, for example an effect like poison. It may
// also be used to modify existing effects, or removing them completely.
type MobEffect struct {
	TargetRuntimeID uint64
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EventID protocol.MobEffectEvent
	// Operation is the operation of the packet. It is either MobEffectAdd, MobEffectModify or MobEffectRemove and
	// specifies the result of the packet client-side.
	EffectID int32
	// EffectType is the ID of the effect to be added, removed or modified. It is one of the constants that may be
	// found above.
	EffectAmplifier int32
	// Particles specifies if viewers of the entity that gets the effect shows particles around it. If set to
	// false, no particles are emitted around the entity.
	ShowParticles bool
	// Duration is the duration of the effect in ticks (20 per second). After the duration has elapsed, the effect
	// will be removed automatically client-side. A negative duration means the effect never expires.
	EffectDurationTicks int32
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
	// Ambient specifies if the effect is ambient. If set to false, it will not get treated as an ambient effect.
	Ambient bool
}

// ID ...
func (*MobEffect) ID() uint32 {
	return IDMobEffect
}

func (pk *MobEffect) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	pk.EventID.Marshal(io)
	io.Varint32(&pk.EffectID)
	io.Varint32(&pk.EffectAmplifier)
	io.Bool(&pk.ShowParticles)
	io.Varint32(&pk.EffectDurationTicks)
	io.PlayerInputTick(&pk.Tick)
	io.Bool(&pk.Ambient)
}
