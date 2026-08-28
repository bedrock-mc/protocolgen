// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EventData interface {
	isEventData()
}

// MarshalEventData reads or writes the EventData union using its canonical wire layout.
func MarshalEventData(io IO, x *EventData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(Achievement)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(Interaction)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(PortalCreated)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(PortalUsed)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(MobKilled)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(CauldronUsed)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(PlayerDied)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(BossKilled)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(SlashCommand)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(MobBorn)
				value.Marshal(io)
				*x = value
			case 10:
				value := new(POICauldronUsed)
				value.Marshal(io)
				*x = value
			case 11:
				value := new(ComposterUsed)
				value.Marshal(io)
				*x = value
			case 12:
				value := new(BellUsed)
				value.Marshal(io)
				*x = value
			case 13:
				value := new(ActorDefinition)
				value.Marshal(io)
				*x = value
			case 14:
				value := new(RaidUpdate)
				value.Marshal(io)
				*x = value
			case 15:
				value := new(TargetBlockHit)
				value.Marshal(io)
				*x = value
			case 16:
				value := new(PiglinBarter)
				value.Marshal(io)
				*x = value
			case 17:
				value := new(PlayerWaxedOrUnwaxedCopper)
				value.Marshal(io)
				*x = value
			case 18:
				value := new(CodeBuilderRuntimeAction)
				value.Marshal(io)
				*x = value
			case 19:
				value := new(CodeBuilderScoreboard)
				value.Marshal(io)
				*x = value
			case 20:
				value := new(ItemUsed)
				value.Marshal(io)
				*x = value
			case 21:
				value := new(Empty)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *Achievement:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *Interaction:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PortalCreated:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PortalUsed:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *MobKilled:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CauldronUsed:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PlayerDied:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BossKilled:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SlashCommand:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *MobBorn:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *POICauldronUsed:
				tag := uint32(10)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ComposterUsed:
				tag := uint32(11)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BellUsed:
				tag := uint32(12)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ActorDefinition:
				tag := uint32(13)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *RaidUpdate:
				tag := uint32(14)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *TargetBlockHit:
				tag := uint32(15)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PiglinBarter:
				tag := uint32(16)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PlayerWaxedOrUnwaxedCopper:
				tag := uint32(17)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CodeBuilderRuntimeAction:
				tag := uint32(18)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CodeBuilderScoreboard:
				tag := uint32(19)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemUsed:
				tag := uint32(20)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *Empty:
				tag := uint32(21)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
