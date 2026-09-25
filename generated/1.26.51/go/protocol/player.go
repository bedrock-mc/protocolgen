// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerActionType int32

const (
	PlayerActionTypeUnknown               PlayerActionType = -1
	PlayerActionTypeStartdestroyblock     PlayerActionType = 0
	PlayerActionTypeAbortdestroyblock     PlayerActionType = 1
	PlayerActionTypeStopdestroyblock      PlayerActionType = 2
	PlayerActionTypeGetupdatedblock       PlayerActionType = 3
	PlayerActionTypeDropitem              PlayerActionType = 4
	PlayerActionTypeStartsleeping         PlayerActionType = 5
	PlayerActionTypeStopsleeping          PlayerActionType = 6
	PlayerActionTypeRespawn               PlayerActionType = 7
	PlayerActionTypeStartjump             PlayerActionType = 8
	PlayerActionTypeStartsprinting        PlayerActionType = 9
	PlayerActionTypeStopsprinting         PlayerActionType = 10
	PlayerActionTypeStartsneaking         PlayerActionType = 11
	PlayerActionTypeStopsneaking          PlayerActionType = 12
	PlayerActionTypeCreativedestroyblock  PlayerActionType = 13
	PlayerActionTypeChangedimensionack    PlayerActionType = 14
	PlayerActionTypeStartgliding          PlayerActionType = 15
	PlayerActionTypeStopgliding           PlayerActionType = 16
	PlayerActionTypeDenydestroyblock      PlayerActionType = 17
	PlayerActionTypeCrackblock            PlayerActionType = 18
	PlayerActionTypeChangeskin            PlayerActionType = 19
	PlayerActionTypeUpdatedenchantingseed PlayerActionType = 20
	PlayerActionTypeStartswimming         PlayerActionType = 21
	PlayerActionTypeStopswimming          PlayerActionType = 22
	PlayerActionTypeStartspinattack       PlayerActionType = 23
	PlayerActionTypeStopspinattack        PlayerActionType = 24
	PlayerActionTypeInteractwithblock     PlayerActionType = 25
	PlayerActionTypePredictdestroyblock   PlayerActionType = 26
	PlayerActionTypeContinuedestroyblock  PlayerActionType = 27
	PlayerActionTypeStartitemuseon        PlayerActionType = 28
	PlayerActionTypeStopitemuseon         PlayerActionType = 29
	PlayerActionTypeHandledteleport       PlayerActionType = 30
	PlayerActionTypeMissedswing           PlayerActionType = 31
	PlayerActionTypeStartcrawling         PlayerActionType = 32
	PlayerActionTypeStopcrawling          PlayerActionType = 33
	PlayerActionTypeStartflying           PlayerActionType = 34
	PlayerActionTypeStopflying            PlayerActionType = 35
	PlayerActionTypeClientackserverdata   PlayerActionType = 36
	PlayerActionTypeStartusingitem        PlayerActionType = 37
	PlayerActionTypeInternalupdate        PlayerActionType = 38
	PlayerActionTypeCount                 PlayerActionType = 39
)

// Marshal reads or writes PlayerActionType through its int32 wire encoding.
func (x *PlayerActionType) Marshal(io IO) { io.Varint32((*int32)(x)) }

// PlayerBlockAction ...
type PlayerBlockActionData struct {
	// Action is the action to be performed, and is one of the constants listed above.
	PlayerActionType PlayerActionType
	// BlockPos is the position of the block that was interacted with.
	Position BlockPos
	// Face is the face of the block that was interacted with.
	Facing int32
}

// Marshal reads or writes PlayerBlockActionData using its canonical wire layout.
func (x *PlayerBlockActionData) Marshal(io IO) {
	x.PlayerActionType.Marshal(io)
	x.Position.Marshal(io)
	io.Varint32(&x.Facing)
}

// PlayerDiedEvent is the event data sent when a player dies.
type PlayerDied struct {
	// AttackerEntityID ...
	InstigatorActorID int32
	// AttackerVariant ...
	InstigatorMobVariant int32
	// EntityDamageCause ...
	DamageSource int32
	// InRaid ...
	DiedInRaid bool
}

func (*PlayerDied) tagEventData() uint32 { return 6 }

// Marshal reads or writes PlayerDied using its canonical wire layout.
func (x *PlayerDied) Marshal(io IO) {
	io.Varint32(&x.InstigatorActorID)
	io.Varint32(&x.InstigatorMobVariant)
	io.Varint32(&x.DamageSource)
	io.Bool(&x.DiedInRaid)
}

type PlayerListData interface {
	Marshaler
	tagPlayerListData() uint32
}

// MarshalPlayerListData reads or writes the PlayerListData union using its canonical wire layout.
func MarshalPlayerListData(io IO, x *PlayerListData) {
	Union(io, x, io.Varuint32, PlayerListData.tagPlayerListData, func(tag uint32) PlayerListData {
		switch tag {
		case 0:
			return new(RemoveEntry)
		case 1:
			return new(AddEntry)
		}
		return nil
	})
}

type PlayerListPacketType uint8

const (
	PlayerListPacketTypeRemove PlayerListPacketType = 1
)

// Marshal reads or writes PlayerListPacketType through its uint8 wire encoding.
func (x *PlayerListPacketType) Marshal(io IO) { io.Uint8((*uint8)(x)) }

type PlayerLocationData interface {
	Marshaler
	tagPlayerLocationData() uint32
}

// MarshalPlayerLocationData reads or writes the PlayerLocationData union using its canonical wire layout.
func MarshalPlayerLocationData(io IO, x *PlayerLocationData) {
	Union(io, x, io.Varuint32, PlayerLocationData.tagPlayerLocationData, func(tag uint32) PlayerLocationData {
		switch tag {
		case 0:
			return new(CoordinatesLocation)
		case 1:
			return new(HiddenLocation)
		}
		return nil
	})
}

type PlayerLocationType int32

const (
	PlayerLocationTypePlayerLocationCoordinates PlayerLocationType = 0
)

// Marshal reads or writes PlayerLocationType through its int32 wire encoding.
func (x *PlayerLocationType) Marshal(io IO) { io.Varint32((*int32)(x)) }

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

// Marshal reads or writes PlayerPermissionLevel through its int8 wire encoding.
func (x *PlayerPermissionLevel) Marshal(io IO) { io.Int8((*int8)(x)) }

type PlayerPositionModeComponentPositionMode uint8

const (
	PlayerPositionModeComponentPositionModeNormal      PlayerPositionModeComponentPositionMode = 0
	PlayerPositionModeComponentPositionModeRespawn     PlayerPositionModeComponentPositionMode = 1
	PlayerPositionModeComponentPositionModeTeleport    PlayerPositionModeComponentPositionMode = 2
	PlayerPositionModeComponentPositionModeOnlyheadrot PlayerPositionModeComponentPositionMode = 3
)

// Marshal reads or writes PlayerPositionModeComponentPositionMode through its uint8 wire encoding.
func (x *PlayerPositionModeComponentPositionMode) Marshal(io IO) { io.Uint8((*uint8)(x)) }

type PlayerRespawnState uint8

const (
	PlayerRespawnStateSearchingforspawn  PlayerRespawnState = 0
	PlayerRespawnStateReadytospawn       PlayerRespawnState = 1
	PlayerRespawnStateClientreadytospawn PlayerRespawnState = 2
)

// Marshal reads or writes PlayerRespawnState through its uint8 wire encoding.
func (x *PlayerRespawnState) Marshal(io IO) { io.Uint8((*uint8)(x)) }

type PlayerScoreboardID struct {
	PlayerUniqueID int64
}

// Marshal reads or writes PlayerScoreboardID using its canonical wire layout.
func (x *PlayerScoreboardID) Marshal(io IO) {
	io.ActorUniqueID(&x.PlayerUniqueID)
}

type PlayerVideoCaptureData interface {
	Marshaler
	tagPlayerVideoCaptureData() uint8
}

// MarshalPlayerVideoCaptureData reads or writes the PlayerVideoCaptureData union using its canonical wire layout.
func MarshalPlayerVideoCaptureData(io IO, x *PlayerVideoCaptureData) {
	Union(io, x, io.Uint8, PlayerVideoCaptureData.tagPlayerVideoCaptureData, func(tag uint8) PlayerVideoCaptureData {
		switch tag {
		case 0:
			return new(StopVideoCapture)
		case 1:
			return new(StartVideoCapture)
		}
		return nil
	})
}

type PlayerWaxedOrUnwaxedCopper struct {
	PlayerWaxedOrUnwaxedCopperBlockID int32
}

func (*PlayerWaxedOrUnwaxedCopper) tagEventData() uint32 { return 17 }

// Marshal reads or writes PlayerWaxedOrUnwaxedCopper using its canonical wire layout.
func (x *PlayerWaxedOrUnwaxedCopper) Marshal(io IO) {
	io.Varint32(&x.PlayerWaxedOrUnwaxedCopperBlockID)
}
