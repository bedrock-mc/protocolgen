// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ActorEventType uint8

const (
	ActorEventTypeNone                             ActorEventType = 0
	ActorEventTypeJump                             ActorEventType = 1
	ActorEventTypeHurt                             ActorEventType = 2
	ActorEventTypeDeath                            ActorEventType = 3
	ActorEventTypeStartAttacking                   ActorEventType = 4
	ActorEventTypeStopAttacking                    ActorEventType = 5
	ActorEventTypeTamingFailed                     ActorEventType = 6
	ActorEventTypeTamingSucceeded                  ActorEventType = 7
	ActorEventTypeShakeWetness                     ActorEventType = 8
	ActorEventTypeEatGrass                         ActorEventType = 10
	ActorEventTypeFishhookBubble                   ActorEventType = 11
	ActorEventTypeFishhookFishPosition             ActorEventType = 12
	ActorEventTypeFishhookHookTime                 ActorEventType = 13
	ActorEventTypeFishhookTease                    ActorEventType = 14
	ActorEventTypeSquidFleeing                     ActorEventType = 15
	ActorEventTypeZombieConverting                 ActorEventType = 16
	ActorEventTypePlayAmbient                      ActorEventType = 17
	ActorEventTypeSpawnAlive                       ActorEventType = 18
	ActorEventTypeStartOfferFlower                 ActorEventType = 19
	ActorEventTypeStopOfferFlower                  ActorEventType = 20
	ActorEventTypeLoveHearts                       ActorEventType = 21
	ActorEventTypeVillagerAngry                    ActorEventType = 22
	ActorEventTypeVillagerHappy                    ActorEventType = 23
	ActorEventTypeWitchHatMagic                    ActorEventType = 24
	ActorEventTypeFireworksExplode                 ActorEventType = 25
	ActorEventTypeInLoveHearts                     ActorEventType = 26
	ActorEventTypeSilverfishMergeAnimation         ActorEventType = 27
	ActorEventTypeGuardianAttackSound              ActorEventType = 28
	ActorEventTypeDrinkPotion                      ActorEventType = 29
	ActorEventTypeThrowPotion                      ActorEventType = 30
	ActorEventTypePrimeTNTCart                     ActorEventType = 31
	ActorEventTypePrimeCreeper                     ActorEventType = 32
	ActorEventTypeAirSupply                        ActorEventType = 33
	ActorEventTypeDeprecatedAddPlayerLevels        ActorEventType = 34
	ActorEventTypeGuardianMiningFatigue            ActorEventType = 35
	ActorEventTypeAgentSwingArm                    ActorEventType = 36
	ActorEventTypeDragonStartDeathAnimation        ActorEventType = 37
	ActorEventTypeGroundDust                       ActorEventType = 38
	ActorEventTypeShake                            ActorEventType = 39
	ActorEventTypeFeed                             ActorEventType = 57
	ActorEventTypeBabyAge                          ActorEventType = 60
	ActorEventTypeInstantDeath                     ActorEventType = 61
	ActorEventTypeNotifyTrade                      ActorEventType = 62
	ActorEventTypeLeashDestroyed                   ActorEventType = 63
	ActorEventTypeCaravanUpdated                   ActorEventType = 64
	ActorEventTypeTalismanActivate                 ActorEventType = 65
	ActorEventTypeDeprecatedUpdateStructureFeature ActorEventType = 66
	ActorEventTypePlayerSpawnedMob                 ActorEventType = 67
	ActorEventTypePuke                             ActorEventType = 68
	ActorEventTypeUpdateStackSize                  ActorEventType = 69
	ActorEventTypeStartSwimming                    ActorEventType = 70
	ActorEventTypeBalloonPop                       ActorEventType = 71
	ActorEventTypeTreasureHunt                     ActorEventType = 72
	ActorEventTypeSummonAgent                      ActorEventType = 73
	ActorEventTypeFinishedChargingItem             ActorEventType = 74
	ActorEventTypeActorGrowUp                      ActorEventType = 76
	ActorEventTypeVibrationDetected                ActorEventType = 77
	ActorEventTypeDrinkMilk                        ActorEventType = 78
	ActorEventTypeShakeWetnessStop                 ActorEventType = 79
	ActorEventTypeKineticDamageDealt               ActorEventType = 80
	ActorEventTypeHurtWithoutReceivingDamage       ActorEventType = 81
)
