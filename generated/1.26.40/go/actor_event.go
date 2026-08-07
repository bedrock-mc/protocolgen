// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type ActorEvent struct {
	TargetRuntimeID ActorRuntimeID
	EventID         ActorEventType
	Data            int32
	FireAtPosition  *mgl32.Vec3
}

func (p *ActorEvent) Encode(w Encoder) error {
	if err := w.Write("ActorEventPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("ActorEventPacket.Event ID", Shape{Kind: "enum", Semantic: "ActorEvent", TypeID: "enums/ActorEvent", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NONE", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "JUMP", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "HURT", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "DEATH", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "START_ATTACKING", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "STOP_ATTACKING", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "TAMING_FAILED", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "TAMING_SUCCEEDED", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "SHAKE_WETNESS", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "EAT_GRASS", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "FISHHOOK_BUBBLE", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "FISHHOOK_FISHPOS", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "FISHHOOK_HOOKTIME", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "FISHHOOK_TEASE", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "SQUID_FLEEING", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "ZOMBIE_CONVERTING", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "PLAY_AMBIENT", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "SPAWN_ALIVE", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "START_OFFER_FLOWER", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "STOP_OFFER_FLOWER", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "LOVE_HEARTS", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "VILLAGER_ANGRY", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "VILLAGER_HAPPY", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "WITCH_HAT_MAGIC", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "FIREWORKS_EXPLODE", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "IN_LOVE_HEARTS", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "SILVERFISH_MERGE_ANIM", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "GUARDIAN_ATTACK_SOUND", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "DRINK_POTION", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "THROW_POTION", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "PRIME_TNTCART", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "PRIME_CREEPER", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "AIR_SUPPLY", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "DEPRECATED_ADD_PLAYER_LEVELS", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "GUARDIAN_MINING_FATIGUE", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "AGENT_SWING_ARM", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "DRAGON_START_DEATH_ANIM", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "GROUND_DUST", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "SHAKE", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "FEED", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "BABY_AGE", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "INSTANT_DEATH", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "NOTIFY_TRADE", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "LEASH_DESTROYED", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "CARAVAN_UPDATED", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "TALISMAN_ACTIVATE", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "DEPRECATED_UPDATE_STRUCTURE_FEATURE", Shape: Shape{Kind: "void"}}, {Value: 67, Name: "PLAYER_SPAWNED_MOB", Shape: Shape{Kind: "void"}}, {Value: 68, Name: "PUKE", Shape: Shape{Kind: "void"}}, {Value: 69, Name: "UPDATE_STACK_SIZE", Shape: Shape{Kind: "void"}}, {Value: 70, Name: "START_SWIMMING", Shape: Shape{Kind: "void"}}, {Value: 71, Name: "BALLOON_POP", Shape: Shape{Kind: "void"}}, {Value: 72, Name: "TREASURE_HUNT", Shape: Shape{Kind: "void"}}, {Value: 73, Name: "SUMMON_AGENT", Shape: Shape{Kind: "void"}}, {Value: 74, Name: "FINISHED_CHARGING_ITEM", Shape: Shape{Kind: "void"}}, {Value: 76, Name: "ACTOR_GROW_UP", Shape: Shape{Kind: "void"}}, {Value: 77, Name: "VIBRATION_DETECTED", Shape: Shape{Kind: "void"}}, {Value: 78, Name: "DRINK_MILK", Shape: Shape{Kind: "void"}}, {Value: 79, Name: "SHAKE_WETNESS_STOP", Shape: Shape{Kind: "void"}}, {Value: 80, Name: "KINETIC_DAMAGE_DEALT", Shape: Shape{Kind: "void"}}, {Value: 81, Name: "HURT_WITHOUT_RECEIVING_DAMAGE", Shape: Shape{Kind: "void"}}}}, p.EventID); err != nil {
		return err
	}
	if err := w.Write("ActorEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Data); err != nil {
		return err
	}
	if err := w.Write("ActorEventPacket.Fire At Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.FireAtPosition); err != nil {
		return err
	}
	return nil
}

func DecodeActorEvent(r Decoder) (ActorEvent, error) {
	var p ActorEvent
	{
		raw, err := r.Read("ActorEventPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field ActorEventPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("ActorEventPacket.Event ID", Shape{Kind: "enum", Semantic: "ActorEvent", TypeID: "enums/ActorEvent", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NONE", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "JUMP", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "HURT", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "DEATH", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "START_ATTACKING", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "STOP_ATTACKING", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "TAMING_FAILED", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "TAMING_SUCCEEDED", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "SHAKE_WETNESS", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "EAT_GRASS", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "FISHHOOK_BUBBLE", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "FISHHOOK_FISHPOS", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "FISHHOOK_HOOKTIME", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "FISHHOOK_TEASE", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "SQUID_FLEEING", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "ZOMBIE_CONVERTING", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "PLAY_AMBIENT", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "SPAWN_ALIVE", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "START_OFFER_FLOWER", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "STOP_OFFER_FLOWER", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "LOVE_HEARTS", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "VILLAGER_ANGRY", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "VILLAGER_HAPPY", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "WITCH_HAT_MAGIC", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "FIREWORKS_EXPLODE", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "IN_LOVE_HEARTS", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "SILVERFISH_MERGE_ANIM", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "GUARDIAN_ATTACK_SOUND", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "DRINK_POTION", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "THROW_POTION", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "PRIME_TNTCART", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "PRIME_CREEPER", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "AIR_SUPPLY", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "DEPRECATED_ADD_PLAYER_LEVELS", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "GUARDIAN_MINING_FATIGUE", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "AGENT_SWING_ARM", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "DRAGON_START_DEATH_ANIM", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "GROUND_DUST", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "SHAKE", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "FEED", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "BABY_AGE", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "INSTANT_DEATH", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "NOTIFY_TRADE", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "LEASH_DESTROYED", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "CARAVAN_UPDATED", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "TALISMAN_ACTIVATE", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "DEPRECATED_UPDATE_STRUCTURE_FEATURE", Shape: Shape{Kind: "void"}}, {Value: 67, Name: "PLAYER_SPAWNED_MOB", Shape: Shape{Kind: "void"}}, {Value: 68, Name: "PUKE", Shape: Shape{Kind: "void"}}, {Value: 69, Name: "UPDATE_STACK_SIZE", Shape: Shape{Kind: "void"}}, {Value: 70, Name: "START_SWIMMING", Shape: Shape{Kind: "void"}}, {Value: 71, Name: "BALLOON_POP", Shape: Shape{Kind: "void"}}, {Value: 72, Name: "TREASURE_HUNT", Shape: Shape{Kind: "void"}}, {Value: 73, Name: "SUMMON_AGENT", Shape: Shape{Kind: "void"}}, {Value: 74, Name: "FINISHED_CHARGING_ITEM", Shape: Shape{Kind: "void"}}, {Value: 76, Name: "ACTOR_GROW_UP", Shape: Shape{Kind: "void"}}, {Value: 77, Name: "VIBRATION_DETECTED", Shape: Shape{Kind: "void"}}, {Value: 78, Name: "DRINK_MILK", Shape: Shape{Kind: "void"}}, {Value: 79, Name: "SHAKE_WETNESS_STOP", Shape: Shape{Kind: "void"}}, {Value: 80, Name: "KINETIC_DAMAGE_DEALT", Shape: Shape{Kind: "void"}}, {Value: 81, Name: "HURT_WITHOUT_RECEIVING_DAMAGE", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorEventType)
		if !ok {
			return p, fmt.Errorf("field ActorEventPacket.Event ID has unexpected decoded type %T", raw)
		}
		p.EventID = value
	}
	{
		raw, err := r.Read("ActorEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ActorEventPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	{
		raw, err := r.Read("ActorEventPacket.Fire At Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field ActorEventPacket.Fire At Position has unexpected decoded type %T", raw)
		}
		p.FireAtPosition = value
	}
	return p, nil
}
