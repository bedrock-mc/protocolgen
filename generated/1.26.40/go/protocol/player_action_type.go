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
