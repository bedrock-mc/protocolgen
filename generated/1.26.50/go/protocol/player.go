// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerActionType int32

const (
	PlayerActionTypeUnknown               PlayerActionType = -1
	PlayerActionTypeStartDestroyBlock     PlayerActionType = 0
	PlayerActionTypeAbortDestroyBlock     PlayerActionType = 1
	PlayerActionTypeStopDestroyBlock      PlayerActionType = 2
	PlayerActionTypeGetUpdatedBlock       PlayerActionType = 3
	PlayerActionTypeDropItem              PlayerActionType = 4
	PlayerActionTypeStartSleeping         PlayerActionType = 5
	PlayerActionTypeStopSleeping          PlayerActionType = 6
	PlayerActionTypeRespawn               PlayerActionType = 7
	PlayerActionTypeStartJump             PlayerActionType = 8
	PlayerActionTypeStartSprinting        PlayerActionType = 9
	PlayerActionTypeStopSprinting         PlayerActionType = 10
	PlayerActionTypeStartSneaking         PlayerActionType = 11
	PlayerActionTypeStopSneaking          PlayerActionType = 12
	PlayerActionTypeCreativeDestroyBlock  PlayerActionType = 13
	PlayerActionTypeChangeDimensionAck    PlayerActionType = 14
	PlayerActionTypeStartGliding          PlayerActionType = 15
	PlayerActionTypeStopGliding           PlayerActionType = 16
	PlayerActionTypeDenyDestroyBlock      PlayerActionType = 17
	PlayerActionTypeCrackBlock            PlayerActionType = 18
	PlayerActionTypeChangeSkin            PlayerActionType = 19
	PlayerActionTypeUpdatedEnchantingSeed PlayerActionType = 20
	PlayerActionTypeStartSwimming         PlayerActionType = 21
	PlayerActionTypeStopSwimming          PlayerActionType = 22
	PlayerActionTypeStartSpinAttack       PlayerActionType = 23
	PlayerActionTypeStopSpinAttack        PlayerActionType = 24
	PlayerActionTypeInteractWithBlock     PlayerActionType = 25
	PlayerActionTypePredictDestroyBlock   PlayerActionType = 26
	PlayerActionTypeContinueDestroyBlock  PlayerActionType = 27
	PlayerActionTypeStartItemUseOn        PlayerActionType = 28
	PlayerActionTypeStopItemUseOn         PlayerActionType = 29
	PlayerActionTypeHandledTeleport       PlayerActionType = 30
	PlayerActionTypeMissedSwing           PlayerActionType = 31
	PlayerActionTypeStartCrawling         PlayerActionType = 32
	PlayerActionTypeStopCrawling          PlayerActionType = 33
	PlayerActionTypeStartFlying           PlayerActionType = 34
	PlayerActionTypeStopFlying            PlayerActionType = 35
	PlayerActionTypeClientAckServerData   PlayerActionType = 36
	PlayerActionTypeStartUsingItem        PlayerActionType = 37
	PlayerActionTypeInternalUpdate        PlayerActionType = 38
	PlayerActionTypeCount                 PlayerActionType = 39
)

// PlayerBlockAction ...
type PlayerBlockActionData struct {
	PlayerActionType PlayerActionType
	Position         BlockPos
	Facing           int32
}

// Marshal reads or writes PlayerBlockActionData using its canonical wire layout.
func (x *PlayerBlockActionData) Marshal(io IO) {
	IntegerFunc(&x.PlayerActionType, io.Varint32)
	x.Position.Marshal(io)
	io.Varint32(&x.Facing)
}

type PlayerDied struct {
	InstigatorActorID    int32
	InstigatorMobVariant int32
	DamageSource         int32
	DiedInRaid           bool
}

func (*PlayerDied) isEventData() {}

// Marshal reads or writes PlayerDied using its canonical wire layout.
func (x *PlayerDied) Marshal(io IO) {
	io.Varint32(&x.InstigatorActorID)
	io.Varint32(&x.InstigatorMobVariant)
	io.Varint32(&x.DamageSource)
	io.Bool(&x.DiedInRaid)
}

type PlayerListData interface {
	isPlayerListData()
}

// MarshalPlayerListData reads or writes the PlayerListData union using its canonical wire layout.
func MarshalPlayerListData(io IO, x *PlayerListData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(RemoveEntry)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(AddEntry)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *RemoveEntry:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AddEntry:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type PlayerListPacketType uint8

const (
	PlayerListPacketTypeRemove PlayerListPacketType = 1
)

type PlayerLocationData interface {
	isPlayerLocationData()
}

// MarshalPlayerLocationData reads or writes the PlayerLocationData union using its canonical wire layout.
func MarshalPlayerLocationData(io IO, x *PlayerLocationData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(CoordinatesLocation)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(HiddenLocation)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *CoordinatesLocation:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *HiddenLocation:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type PlayerLocationType int32

const (
	PlayerLocationTypePlayerLocationCoordinates PlayerLocationType = 0
)

type PlayerPartyInfo struct {
	PartyID       string
	IsPartyLeader bool
}

// Marshal reads or writes PlayerPartyInfo using its canonical wire layout.
func (x *PlayerPartyInfo) Marshal(io IO) {
	io.StringLimits(&x.PartyID, 0, 49)
	io.Bool(&x.IsPartyLeader)
}

type PlayerPermissionLevel int8

const (
	PlayerPermissionLevelVisitor  PlayerPermissionLevel = 0
	PlayerPermissionLevelMember   PlayerPermissionLevel = 1
	PlayerPermissionLevelOperator PlayerPermissionLevel = 2
	PlayerPermissionLevelCustom   PlayerPermissionLevel = 3
)

type PlayerPositionModeComponentPositionMode uint8

const (
	PlayerPositionModeComponentPositionModeNormal      PlayerPositionModeComponentPositionMode = 0
	PlayerPositionModeComponentPositionModeRespawn     PlayerPositionModeComponentPositionMode = 1
	PlayerPositionModeComponentPositionModeTeleport    PlayerPositionModeComponentPositionMode = 2
	PlayerPositionModeComponentPositionModeOnlyHeadRot PlayerPositionModeComponentPositionMode = 3
)

type PlayerRespawnState uint8

const (
	PlayerRespawnStateSearchingForSpawn  PlayerRespawnState = 0
	PlayerRespawnStateReadyToSpawn       PlayerRespawnState = 1
	PlayerRespawnStateClientReadyToSpawn PlayerRespawnState = 2
)

type PlayerScoreboardID struct {
	PlayerUniqueID int64
}

// Marshal reads or writes PlayerScoreboardID using its canonical wire layout.
func (x *PlayerScoreboardID) Marshal(io IO) {
	io.Varint64(&x.PlayerUniqueID)
}

type PlayerVideoCaptureData interface {
	isPlayerVideoCaptureData()
}

// MarshalPlayerVideoCaptureData reads or writes the PlayerVideoCaptureData union using its canonical wire layout.
func MarshalPlayerVideoCaptureData(io IO, x *PlayerVideoCaptureData) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(StopVideoCapture)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(StartVideoCapture)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *StopVideoCapture:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *StartVideoCapture:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type PlayerWaxedOrUnwaxedCopper struct {
	PlayerWaxedOrUnwaxedCopperBlockID int32
}

func (*PlayerWaxedOrUnwaxedCopper) isEventData() {}

// Marshal reads or writes PlayerWaxedOrUnwaxedCopper using its canonical wire layout.
func (x *PlayerWaxedOrUnwaxedCopper) Marshal(io IO) {
	io.Varint32(&x.PlayerWaxedOrUnwaxedCopperBlockID)
}
