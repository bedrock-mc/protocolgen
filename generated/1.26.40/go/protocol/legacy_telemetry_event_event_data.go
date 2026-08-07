// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryEventEventData interface {
	isLegacyTelemetryEventEventData()
}

// MarshalLegacyTelemetryEventEventData reads or writes the LegacyTelemetryEventEventData union using its canonical wire layout.
func MarshalLegacyTelemetryEventEventData(io IO, x *LegacyTelemetryEventEventData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value LegacyTelemetryEventAchievement
				value.Marshal(io)
				*x = value
			case 1:
				var value LegacyTelemetryEventInteraction
				value.Marshal(io)
				*x = value
			case 2:
				var value LegacyTelemetryEventPortalCreated
				value.Marshal(io)
				*x = value
			case 3:
				var value LegacyTelemetryEventPortalUsed
				value.Marshal(io)
				*x = value
			case 4:
				var value LegacyTelemetryEventMobKilled
				value.Marshal(io)
				*x = value
			case 5:
				var value LegacyTelemetryEventCauldronUsed
				value.Marshal(io)
				*x = value
			case 6:
				var value LegacyTelemetryEventPlayerDied
				value.Marshal(io)
				*x = value
			case 7:
				var value LegacyTelemetryEventBossKilled
				value.Marshal(io)
				*x = value
			case 8:
				var value LegacyTelemetryEventSlashCommand
				value.Marshal(io)
				*x = value
			case 9:
				var value LegacyTelemetryEventMobBorn
				value.Marshal(io)
				*x = value
			case 10:
				var value LegacyTelemetryEventPOICauldronUsed
				value.Marshal(io)
				*x = value
			case 11:
				var value LegacyTelemetryEventComposterUsed
				value.Marshal(io)
				*x = value
			case 12:
				var value LegacyTelemetryEventBellUsed
				value.Marshal(io)
				*x = value
			case 13:
				var value LegacyTelemetryEventActorDefinition
				value.Marshal(io)
				*x = value
			case 14:
				var value LegacyTelemetryEventRaidUpdate
				value.Marshal(io)
				*x = value
			case 15:
				var value LegacyTelemetryEventTargetBlockHit
				value.Marshal(io)
				*x = value
			case 16:
				var value LegacyTelemetryEventPiglinBarter
				value.Marshal(io)
				*x = value
			case 17:
				var value LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper
				value.Marshal(io)
				*x = value
			case 18:
				var value LegacyTelemetryEventCodeBuilderRuntimeAction
				value.Marshal(io)
				*x = value
			case 19:
				var value LegacyTelemetryEventCodeBuilderScoreboard
				value.Marshal(io)
				*x = value
			case 20:
				var value LegacyTelemetryEventItemUsed
				value.Marshal(io)
				*x = value
			case 21:
				var value LegacyTelemetryEventEmpty
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case LegacyTelemetryEventAchievement:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventInteraction:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPortalCreated:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPortalUsed:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventMobKilled:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventCauldronUsed:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPlayerDied:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventBossKilled:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventSlashCommand:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventMobBorn:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPOICauldronUsed:
				tag := uint32(10)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventComposterUsed:
				tag := uint32(11)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventBellUsed:
				tag := uint32(12)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventActorDefinition:
				tag := uint32(13)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventRaidUpdate:
				tag := uint32(14)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventTargetBlockHit:
				tag := uint32(15)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPiglinBarter:
				tag := uint32(16)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper:
				tag := uint32(17)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventCodeBuilderRuntimeAction:
				tag := uint32(18)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventCodeBuilderScoreboard:
				tag := uint32(19)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventItemUsed:
				tag := uint32(20)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LegacyTelemetryEventEmpty:
				tag := uint32(21)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
