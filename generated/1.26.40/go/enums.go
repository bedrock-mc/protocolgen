// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ActorEventType uint8

const (
	ActorEventTypeNONE                             ActorEventType = 0
	ActorEventTypeJUMP                             ActorEventType = 1
	ActorEventTypeHURT                             ActorEventType = 2
	ActorEventTypeDEATH                            ActorEventType = 3
	ActorEventTypeSTARTATTACKING                   ActorEventType = 4
	ActorEventTypeSTOPATTACKING                    ActorEventType = 5
	ActorEventTypeTAMINGFAILED                     ActorEventType = 6
	ActorEventTypeTAMINGSUCCEEDED                  ActorEventType = 7
	ActorEventTypeSHAKEWETNESS                     ActorEventType = 8
	ActorEventTypeEATGRASS                         ActorEventType = 10
	ActorEventTypeFISHHOOKBUBBLE                   ActorEventType = 11
	ActorEventTypeFISHHOOKFISHPOS                  ActorEventType = 12
	ActorEventTypeFISHHOOKHOOKTIME                 ActorEventType = 13
	ActorEventTypeFISHHOOKTEASE                    ActorEventType = 14
	ActorEventTypeSQUIDFLEEING                     ActorEventType = 15
	ActorEventTypeZOMBIECONVERTING                 ActorEventType = 16
	ActorEventTypePLAYAMBIENT                      ActorEventType = 17
	ActorEventTypeSPAWNALIVE                       ActorEventType = 18
	ActorEventTypeSTARTOFFERFLOWER                 ActorEventType = 19
	ActorEventTypeSTOPOFFERFLOWER                  ActorEventType = 20
	ActorEventTypeLOVEHEARTS                       ActorEventType = 21
	ActorEventTypeVILLAGERANGRY                    ActorEventType = 22
	ActorEventTypeVILLAGERHAPPY                    ActorEventType = 23
	ActorEventTypeWITCHHATMAGIC                    ActorEventType = 24
	ActorEventTypeFIREWORKSEXPLODE                 ActorEventType = 25
	ActorEventTypeINLOVEHEARTS                     ActorEventType = 26
	ActorEventTypeSILVERFISHMERGEANIM              ActorEventType = 27
	ActorEventTypeGUARDIANATTACKSOUND              ActorEventType = 28
	ActorEventTypeDRINKPOTION                      ActorEventType = 29
	ActorEventTypeTHROWPOTION                      ActorEventType = 30
	ActorEventTypePRIMETNTCART                     ActorEventType = 31
	ActorEventTypePRIMECREEPER                     ActorEventType = 32
	ActorEventTypeAIRSUPPLY                        ActorEventType = 33
	ActorEventTypeDEPRECATEDADDPLAYERLEVELS        ActorEventType = 34
	ActorEventTypeGUARDIANMININGFATIGUE            ActorEventType = 35
	ActorEventTypeAGENTSWINGARM                    ActorEventType = 36
	ActorEventTypeDRAGONSTARTDEATHANIM             ActorEventType = 37
	ActorEventTypeGROUNDDUST                       ActorEventType = 38
	ActorEventTypeSHAKE                            ActorEventType = 39
	ActorEventTypeFEED                             ActorEventType = 57
	ActorEventTypeBABYAGE                          ActorEventType = 60
	ActorEventTypeINSTANTDEATH                     ActorEventType = 61
	ActorEventTypeNOTIFYTRADE                      ActorEventType = 62
	ActorEventTypeLEASHDESTROYED                   ActorEventType = 63
	ActorEventTypeCARAVANUPDATED                   ActorEventType = 64
	ActorEventTypeTALISMANACTIVATE                 ActorEventType = 65
	ActorEventTypeDEPRECATEDUPDATESTRUCTUREFEATURE ActorEventType = 66
	ActorEventTypePLAYERSPAWNEDMOB                 ActorEventType = 67
	ActorEventTypePUKE                             ActorEventType = 68
	ActorEventTypeUPDATESTACKSIZE                  ActorEventType = 69
	ActorEventTypeSTARTSWIMMING                    ActorEventType = 70
	ActorEventTypeBALLOONPOP                       ActorEventType = 71
	ActorEventTypeTREASUREHUNT                     ActorEventType = 72
	ActorEventTypeSUMMONAGENT                      ActorEventType = 73
	ActorEventTypeFINISHEDCHARGINGITEM             ActorEventType = 74
	ActorEventTypeACTORGROWUP                      ActorEventType = 76
	ActorEventTypeVIBRATIONDETECTED                ActorEventType = 77
	ActorEventTypeDRINKMILK                        ActorEventType = 78
	ActorEventTypeSHAKEWETNESSSTOP                 ActorEventType = 79
	ActorEventTypeKINETICDAMAGEDEALT               ActorEventType = 80
	ActorEventTypeHURTWITHOUTRECEIVINGDAMAGE       ActorEventType = 81
)

type ActorLinkType uint8

const (
	ActorLinkTypeNone      ActorLinkType = 0
	ActorLinkTypeRiding    ActorLinkType = 1
	ActorLinkTypePassenger ActorLinkType = 2
)

type ActorType int32

const (
	ActorTypeUndefined                  ActorType = 1
	ActorTypeItemEntity                 ActorType = 64
	ActorTypePrimedTnt                  ActorType = 65
	ActorTypeFallingBlock               ActorType = 66
	ActorTypeMovingBlock                ActorType = 67
	ActorTypeExperience                 ActorType = 69
	ActorTypeEyeOfEnder                 ActorType = 70
	ActorTypeEnderCrystal               ActorType = 71
	ActorTypeFireworksRocket            ActorType = 72
	ActorTypeFishingHook                ActorType = 77
	ActorTypeChalkboard                 ActorType = 78
	ActorTypePainting                   ActorType = 83
	ActorTypeLeashKnot                  ActorType = 88
	ActorTypeBoatRideable               ActorType = 90
	ActorTypeLightningBolt              ActorType = 93
	ActorTypeAreaEffectCloud            ActorType = 95
	ActorTypeBalloon                    ActorType = 107
	ActorTypeShield                     ActorType = 117
	ActorTypeLectern                    ActorType = 119
	ActorTypeOminousItemSpawner         ActorType = 145
	ActorTypeCushion                    ActorType = 154
	ActorTypeChestBoatRideable          ActorType = 218
	ActorTypeMob                        ActorType = 256
	ActorTypeNpc                        ActorType = 307
	ActorTypeAgent                      ActorType = 312
	ActorTypeArmorStand                 ActorType = 317
	ActorTypeTripodCamera               ActorType = 318
	ActorTypePlayer                     ActorType = 319
	ActorTypeBee                        ActorType = 378
	ActorTypePiglin                     ActorType = 379
	ActorTypePiglinBrute                ActorType = 383
	ActorTypeAllay                      ActorType = 390
	ActorTypePathfinderMob              ActorType = 768
	ActorTypeIronGolem                  ActorType = 788
	ActorTypeSnowGolem                  ActorType = 789
	ActorTypeWanderingTrader            ActorType = 886
	ActorTypeCopperGolem                ActorType = 916
	ActorTypeSulfurCube                 ActorType = 921
	ActorTypeMonster                    ActorType = 2816
	ActorTypeCreeper                    ActorType = 2849
	ActorTypeSlime                      ActorType = 2853
	ActorTypeEnderMan                   ActorType = 2854
	ActorTypeGhast                      ActorType = 2857
	ActorTypeLavaSlime                  ActorType = 2858
	ActorTypeBlaze                      ActorType = 2859
	ActorTypeWitch                      ActorType = 2861
	ActorTypeGuardian                   ActorType = 2865
	ActorTypeElderGuardian              ActorType = 2866
	ActorTypeDragon                     ActorType = 2869
	ActorTypeShulker                    ActorType = 2870
	ActorTypeVindicator                 ActorType = 2873
	ActorTypeIllagerBeast               ActorType = 2875
	ActorTypeEvocationIllager           ActorType = 2920
	ActorTypeVex                        ActorType = 2921
	ActorTypePillager                   ActorType = 2930
	ActorTypeElderGuardianGhost         ActorType = 2936
	ActorTypeWarden                     ActorType = 2947
	ActorTypeBreeze                     ActorType = 2956
	ActorTypeCreaking                   ActorType = 2962
	ActorTypeAnimal                     ActorType = 4864
	ActorTypeChicken                    ActorType = 4874
	ActorTypeCow                        ActorType = 4875
	ActorTypePig                        ActorType = 4876
	ActorTypeSheep                      ActorType = 4877
	ActorTypeMushroomCow                ActorType = 4880
	ActorTypeRabbit                     ActorType = 4882
	ActorTypePolarBear                  ActorType = 4892
	ActorTypeLlama                      ActorType = 4893
	ActorTypeTurtle                     ActorType = 4938
	ActorTypePanda                      ActorType = 4977
	ActorTypeFox                        ActorType = 4985
	ActorTypeHoglin                     ActorType = 4988
	ActorTypeStrider                    ActorType = 4989
	ActorTypeGoat                       ActorType = 4992
	ActorTypeAxolotl                    ActorType = 4994
	ActorTypeFrog                       ActorType = 4996
	ActorTypeCamel                      ActorType = 5002
	ActorTypeSniffer                    ActorType = 5003
	ActorTypeArmadillo                  ActorType = 5006
	ActorTypeHappyGhast                 ActorType = 5011
	ActorTypeTraderLlama                ActorType = 5021
	ActorTypeWaterAnimal                ActorType = 8960
	ActorTypeSquid                      ActorType = 8977
	ActorTypeDolphin                    ActorType = 8991
	ActorTypePufferfish                 ActorType = 9068
	ActorTypeSalmon                     ActorType = 9069
	ActorTypeTropicalfish               ActorType = 9071
	ActorTypeFish                       ActorType = 9072
	ActorTypeGlowSquid                  ActorType = 9089
	ActorTypeTadpole                    ActorType = 9093
	ActorTypeNautilus                   ActorType = 9109
	ActorTypeTamableAnimal              ActorType = 21248
	ActorTypeWolf                       ActorType = 21262
	ActorTypeOcelot                     ActorType = 21270
	ActorTypeParrot                     ActorType = 21278
	ActorTypeCat                        ActorType = 21323
	ActorTypeAmbient                    ActorType = 33024
	ActorTypeBat                        ActorType = 33043
	ActorTypeUndeadMonster              ActorType = 68352
	ActorTypePigZombie                  ActorType = 68388
	ActorTypeWitherBoss                 ActorType = 68404
	ActorTypePhantom                    ActorType = 68410
	ActorTypeZoglin                     ActorType = 68478
	ActorTypeCamelHusk                  ActorType = 70552
	ActorTypeZombieNautilus             ActorType = 74646
	ActorTypeZombieMonster              ActorType = 199424
	ActorTypeZombie                     ActorType = 199456
	ActorTypeZombieVillager             ActorType = 199468
	ActorTypeHusk                       ActorType = 199471
	ActorTypeDrowned                    ActorType = 199534
	ActorTypeZombieVillagerV2           ActorType = 199540
	ActorTypeArthropod                  ActorType = 264960
	ActorTypeSpider                     ActorType = 264995
	ActorTypeSilverfish                 ActorType = 264999
	ActorTypeCaveSpider                 ActorType = 265000
	ActorTypeEndermite                  ActorType = 265015
	ActorTypeMinecart                   ActorType = 524288
	ActorTypeMinecartRideable           ActorType = 524372
	ActorTypeMinecartHopper             ActorType = 524384
	ActorTypeMinecartTNT                ActorType = 524385
	ActorTypeMinecartChest              ActorType = 524386
	ActorTypeMinecartFurnace            ActorType = 524387
	ActorTypeMinecartCommandBlock       ActorType = 524388
	ActorTypeSkeletonMonster            ActorType = 1116928
	ActorTypeSkeleton                   ActorType = 1116962
	ActorTypeStray                      ActorType = 1116974
	ActorTypeWitherSkeleton             ActorType = 1116976
	ActorTypeBogged                     ActorType = 1117072
	ActorTypeParched                    ActorType = 1117079
	ActorTypeEquineAnimal               ActorType = 2118400
	ActorTypeHorse                      ActorType = 2118423
	ActorTypeDonkey                     ActorType = 2118424
	ActorTypeMule                       ActorType = 2118425
	ActorTypeSkeletonHorse              ActorType = 2183962
	ActorTypeZombieHorse                ActorType = 2183963
	ActorTypeProjectile                 ActorType = 4194304
	ActorTypeExperiencePotion           ActorType = 4194372
	ActorTypeShulkerBullet              ActorType = 4194380
	ActorTypeDragonFireball             ActorType = 4194383
	ActorTypeSnowball                   ActorType = 4194385
	ActorTypeThrownEgg                  ActorType = 4194386
	ActorTypeLargeFireball              ActorType = 4194389
	ActorTypeThrownPotion               ActorType = 4194390
	ActorTypeEnderpearl                 ActorType = 4194391
	ActorTypeWitherSkull                ActorType = 4194393
	ActorTypeWitherSkullDangerous       ActorType = 4194395
	ActorTypeSmallFireball              ActorType = 4194398
	ActorTypeLingeringPotion            ActorType = 4194405
	ActorTypeLlamaSpit                  ActorType = 4194406
	ActorTypeEvocationFang              ActorType = 4194407
	ActorTypeIceBomb                    ActorType = 4194410
	ActorTypeBreezeWindChargeProjectile ActorType = 4194445
	ActorTypeWindChargeProjectile       ActorType = 4194447
	ActorTypeAbstractArrow              ActorType = 8388608
	ActorTypeTrident                    ActorType = 12582985
	ActorTypeArrow                      ActorType = 12582992
	ActorTypeVillagerBase               ActorType = 16777984
	ActorTypeVillager                   ActorType = 16777999
	ActorTypeVillagerV2                 ActorType = 16778099
)

type AgentActionType int32

const (
	AgentActionTypeAttack            AgentActionType = 1
	AgentActionTypeCollect           AgentActionType = 2
	AgentActionTypeDestroy           AgentActionType = 3
	AgentActionTypeDetectRedstone    AgentActionType = 4
	AgentActionTypeDetectObstacle    AgentActionType = 5
	AgentActionTypeDrop              AgentActionType = 6
	AgentActionTypeDropAll           AgentActionType = 7
	AgentActionTypeInspect           AgentActionType = 8
	AgentActionTypeInspectData       AgentActionType = 9
	AgentActionTypeInspectItemCount  AgentActionType = 10
	AgentActionTypeInspectItemDetail AgentActionType = 11
	AgentActionTypeInspectItemSpace  AgentActionType = 12
	AgentActionTypeInteract          AgentActionType = 13
	AgentActionTypeMove              AgentActionType = 14
	AgentActionTypePlaceBlock        AgentActionType = 15
	AgentActionTypeTill              AgentActionType = 16
	AgentActionTypeTransferItemTo    AgentActionType = 17
	AgentActionTypeTurn              AgentActionType = 18
)

type AgentAnimationType uint8

const (
	AgentAnimationTypeArmSwing AgentAnimationType = 0
	AgentAnimationTypeShrug    AgentAnimationType = 1
)

type AnimateAction uint8

const (
	AnimateActionNoAction         AnimateAction = 0
	AnimateActionSwing            AnimateAction = 1
	AnimateActionWakeUp           AnimateAction = 3
	AnimateActionCriticalHit      AnimateAction = 4
	AnimateActionMagicCriticalHit AnimateAction = 5
)

type AnimationMode uint8

const (
	AnimationModeNone   AnimationMode = 0
	AnimationModeLayers AnimationMode = 1
	AnimationModeBlocks AnimationMode = 2
)

type BossBarColor uint8

const (
	BossBarColorPINK          BossBarColor = 0
	BossBarColorBLUE          BossBarColor = 1
	BossBarColorRED           BossBarColor = 2
	BossBarColorGREEN         BossBarColor = 3
	BossBarColorYELLOW        BossBarColor = 4
	BossBarColorPURPLE        BossBarColor = 5
	BossBarColorREBECCAPURPLE BossBarColor = 6
	BossBarColorWHITE         BossBarColor = 7
)

type BossBarOverlay uint8

const (
	BossBarOverlayPROGRESS  BossBarOverlay = 0
	BossBarOverlayNOTCHED6  BossBarOverlay = 1
	BossBarOverlayNOTCHED10 BossBarOverlay = 2
	BossBarOverlayNOTCHED12 BossBarOverlay = 3
	BossBarOverlayNOTCHED20 BossBarOverlay = 4
)

type BossEventUpdateType uint8

const (
	BossEventUpdateTypeAdd              BossEventUpdateType = 0
	BossEventUpdateTypePlayerAdded      BossEventUpdateType = 1
	BossEventUpdateTypeRemove           BossEventUpdateType = 2
	BossEventUpdateTypePlayerRemoved    BossEventUpdateType = 3
	BossEventUpdateTypeUpdatePercent    BossEventUpdateType = 4
	BossEventUpdateTypeUpdateName       BossEventUpdateType = 5
	BossEventUpdateTypeUpdateProperties BossEventUpdateType = 6
	BossEventUpdateTypeUpdateStyle      BossEventUpdateType = 7
	BossEventUpdateTypeQuery            BossEventUpdateType = 8
)

type BuildPlatform int32

const (
	BuildPlatformUnknown      BuildPlatform = -1
	BuildPlatformGoogle       BuildPlatform = 1
	BuildPlatformIOS          BuildPlatform = 2
	BuildPlatformOSX          BuildPlatform = 3
	BuildPlatformAmazon       BuildPlatform = 4
	BuildPlatformGearVR       BuildPlatform = 5
	BuildPlatformUWP          BuildPlatform = 7
	BuildPlatformWin32        BuildPlatform = 8
	BuildPlatformDedicated    BuildPlatform = 9
	BuildPlatformTvOS         BuildPlatform = 10
	BuildPlatformSony         BuildPlatform = 11
	BuildPlatformNx           BuildPlatform = 12
	BuildPlatformXbox         BuildPlatform = 13
	BuildPlatformWindowsPhone BuildPlatform = 14
	BuildPlatformLinux        BuildPlatform = 15
)

type CameraAimAssistAction uint8

const (
	CameraAimAssistActionSet   CameraAimAssistAction = 0
	CameraAimAssistActionClear CameraAimAssistAction = 1
)

type CameraAimAssistPresetsPacketOperation uint8

const (
	CameraAimAssistPresetsPacketOperationSet           CameraAimAssistPresetsPacketOperation = 0
	CameraAimAssistPresetsPacketOperationAddToExisting CameraAimAssistPresetsPacketOperation = 1
)

type CameraAimAssistTargetMode uint8

const (
	CameraAimAssistTargetModeAngle    CameraAimAssistTargetMode = 0
	CameraAimAssistTargetModeDistance CameraAimAssistTargetMode = 1
)

type CameraAimAssistTargetModeType uint8

const (
	CameraAimAssistTargetModeTypeAngle    CameraAimAssistTargetModeType = 0
	CameraAimAssistTargetModeTypeDistance CameraAimAssistTargetModeType = 1
)

type CameraPresetAudioListener uint8

const (
	CameraPresetAudioListenerCamera CameraPresetAudioListener = 0
	CameraPresetAudioListenerPlayer CameraPresetAudioListener = 1
)

type CameraShakeAction uint8

const (
	CameraShakeActionAdd  CameraShakeAction = 0
	CameraShakeActionStop CameraShakeAction = 1
)

type CameraShakeType uint8

const (
	CameraShakeTypePositional CameraShakeType = 0
	CameraShakeTypeRotational CameraShakeType = 1
)

type ChatRestrictionLevel uint8

const (
	ChatRestrictionLevelNone     ChatRestrictionLevel = 0
	ChatRestrictionLevelDropped  ChatRestrictionLevel = 1
	ChatRestrictionLevelDisabled ChatRestrictionLevel = 2
)

type ClientCameraAimAssistPacketAction uint8

const (
	ClientCameraAimAssistPacketActionSetFromCameraPreset ClientCameraAimAssistPacketAction = 0
	ClientCameraAimAssistPacketActionClear               ClientCameraAimAssistPacketAction = 1
)

type ClientPlayMode uint32

const (
	ClientPlayModeNormal              ClientPlayMode = 0
	ClientPlayModeTeaser              ClientPlayMode = 1
	ClientPlayModeScreen              ClientPlayMode = 2
	ClientPlayModeViewer              ClientPlayMode = 3
	ClientPlayModeReality             ClientPlayMode = 4
	ClientPlayModePlacement           ClientPlayMode = 5
	ClientPlayModeLivingRoom          ClientPlayMode = 6
	ClientPlayModeExitLevel           ClientPlayMode = 7
	ClientPlayModeExitLevelLivingRoom ClientPlayMode = 8
	ClientPlayModeNumModes            ClientPlayMode = 9
)

type ClientboundTextureShiftAction uint8

const (
	ClientboundTextureShiftActionInvalid    ClientboundTextureShiftAction = 0
	ClientboundTextureShiftActionInitialize ClientboundTextureShiftAction = 1
	ClientboundTextureShiftActionStart      ClientboundTextureShiftAction = 2
	ClientboundTextureShiftActionSetEnabled ClientboundTextureShiftAction = 3
	ClientboundTextureShiftActionSync       ClientboundTextureShiftAction = 4
)

type CodeBuilderExecutionStateCodeStatus uint8

const (
	CodeBuilderExecutionStateCodeStatusNone       CodeBuilderExecutionStateCodeStatus = 0
	CodeBuilderExecutionStateCodeStatusNotStarted CodeBuilderExecutionStateCodeStatus = 1
	CodeBuilderExecutionStateCodeStatusInProgress CodeBuilderExecutionStateCodeStatus = 2
	CodeBuilderExecutionStateCodeStatusPaused     CodeBuilderExecutionStateCodeStatus = 3
	CodeBuilderExecutionStateCodeStatusError      CodeBuilderExecutionStateCodeStatus = 4
	CodeBuilderExecutionStateCodeStatusSucceeded  CodeBuilderExecutionStateCodeStatus = 5
)

type CodeBuilderStorageQueryOptionsCategory uint8

const (
	CodeBuilderStorageQueryOptionsCategoryNone          CodeBuilderStorageQueryOptionsCategory = 0
	CodeBuilderStorageQueryOptionsCategoryCodeStatus    CodeBuilderStorageQueryOptionsCategory = 1
	CodeBuilderStorageQueryOptionsCategoryInstantiation CodeBuilderStorageQueryOptionsCategory = 2
)

type CodeBuilderStorageQueryOptionsOperation uint8

const (
	CodeBuilderStorageQueryOptionsOperationNone  CodeBuilderStorageQueryOptionsOperation = 0
	CodeBuilderStorageQueryOptionsOperationGet   CodeBuilderStorageQueryOptionsOperation = 1
	CodeBuilderStorageQueryOptionsOperationSet   CodeBuilderStorageQueryOptionsOperation = 2
	CodeBuilderStorageQueryOptionsOperationReset CodeBuilderStorageQueryOptionsOperation = 3
)

type CommandPermissionLevel uint8

const (
	CommandPermissionLevelAny           CommandPermissionLevel = 0
	CommandPermissionLevelGameDirectors CommandPermissionLevel = 1
	CommandPermissionLevelAdmin         CommandPermissionLevel = 2
	CommandPermissionLevelHost          CommandPermissionLevel = 3
	CommandPermissionLevelOwner         CommandPermissionLevel = 4
	CommandPermissionLevelInternal      CommandPermissionLevel = 5
)

type ConnectionDisconnectFailReason int32

const (
	ConnectionDisconnectFailReasonUnknown                                       ConnectionDisconnectFailReason = 0
	ConnectionDisconnectFailReasonCantConnectNoInternet                         ConnectionDisconnectFailReason = 1
	ConnectionDisconnectFailReasonNoPermissions                                 ConnectionDisconnectFailReason = 2
	ConnectionDisconnectFailReasonUnrecoverableError                            ConnectionDisconnectFailReason = 3
	ConnectionDisconnectFailReasonThirdPartyBlocked                             ConnectionDisconnectFailReason = 4
	ConnectionDisconnectFailReasonThirdPartyNoInternet                          ConnectionDisconnectFailReason = 5
	ConnectionDisconnectFailReasonThirdPartyBadIP                               ConnectionDisconnectFailReason = 6
	ConnectionDisconnectFailReasonThirdPartyNoServerOrServerLocked              ConnectionDisconnectFailReason = 7
	ConnectionDisconnectFailReasonVersionMismatch                               ConnectionDisconnectFailReason = 8
	ConnectionDisconnectFailReasonSkinIssue                                     ConnectionDisconnectFailReason = 9
	ConnectionDisconnectFailReasonInviteSessionNotFound                         ConnectionDisconnectFailReason = 10
	ConnectionDisconnectFailReasonEduLevelSettingsMissing                       ConnectionDisconnectFailReason = 11
	ConnectionDisconnectFailReasonLocalServerNotFound                           ConnectionDisconnectFailReason = 12
	ConnectionDisconnectFailReasonLegacyDisconnect                              ConnectionDisconnectFailReason = 13
	ConnectionDisconnectFailReasonINTERNALUserLeaveGameAttempted                ConnectionDisconnectFailReason = 14
	ConnectionDisconnectFailReasonPlatformLockedSkinsError                      ConnectionDisconnectFailReason = 15
	ConnectionDisconnectFailReasonRealmsWorldUnassigned                         ConnectionDisconnectFailReason = 16
	ConnectionDisconnectFailReasonRealmsServerCantConnect                       ConnectionDisconnectFailReason = 17
	ConnectionDisconnectFailReasonRealmsServerHidden                            ConnectionDisconnectFailReason = 18
	ConnectionDisconnectFailReasonRealmsServerDisabledBeta                      ConnectionDisconnectFailReason = 19
	ConnectionDisconnectFailReasonRealmsServerDisabled                          ConnectionDisconnectFailReason = 20
	ConnectionDisconnectFailReasonCrossPlatformDisabled                         ConnectionDisconnectFailReason = 21
	ConnectionDisconnectFailReasonTESTONLYCantConnect                           ConnectionDisconnectFailReason = 22
	ConnectionDisconnectFailReasonSessionNotFound                               ConnectionDisconnectFailReason = 23
	ConnectionDisconnectFailReasonClientSettingsIncompatibleWithServer          ConnectionDisconnectFailReason = 24
	ConnectionDisconnectFailReasonServerFull                                    ConnectionDisconnectFailReason = 25
	ConnectionDisconnectFailReasonInvalidPlatformSkin                           ConnectionDisconnectFailReason = 26
	ConnectionDisconnectFailReasonEditionVersionMismatch                        ConnectionDisconnectFailReason = 27
	ConnectionDisconnectFailReasonEditionMismatch                               ConnectionDisconnectFailReason = 28
	ConnectionDisconnectFailReasonLevelNewerThanExeVersion                      ConnectionDisconnectFailReason = 29
	ConnectionDisconnectFailReasonINTERNALNoFailOccurred                        ConnectionDisconnectFailReason = 30
	ConnectionDisconnectFailReasonBannedSkin                                    ConnectionDisconnectFailReason = 31
	ConnectionDisconnectFailReasonTimeout                                       ConnectionDisconnectFailReason = 32
	ConnectionDisconnectFailReasonServerNotFound                                ConnectionDisconnectFailReason = 33
	ConnectionDisconnectFailReasonOutdatedServer                                ConnectionDisconnectFailReason = 34
	ConnectionDisconnectFailReasonOutdatedClient                                ConnectionDisconnectFailReason = 35
	ConnectionDisconnectFailReasonNoPremiumPlatform                             ConnectionDisconnectFailReason = 36
	ConnectionDisconnectFailReasonMultiplayerDisabled                           ConnectionDisconnectFailReason = 37
	ConnectionDisconnectFailReasonNoWiFi                                        ConnectionDisconnectFailReason = 38
	ConnectionDisconnectFailReasonWorldCorruption                               ConnectionDisconnectFailReason = 39
	ConnectionDisconnectFailReasonNoReason                                      ConnectionDisconnectFailReason = 40
	ConnectionDisconnectFailReasonDisconnected                                  ConnectionDisconnectFailReason = 41
	ConnectionDisconnectFailReasonInvalidPlayer                                 ConnectionDisconnectFailReason = 42
	ConnectionDisconnectFailReasonLoggedInOtherLocation                         ConnectionDisconnectFailReason = 43
	ConnectionDisconnectFailReasonServerIdConflict                              ConnectionDisconnectFailReason = 44
	ConnectionDisconnectFailReasonNotAllowed                                    ConnectionDisconnectFailReason = 45
	ConnectionDisconnectFailReasonNotAuthenticated                              ConnectionDisconnectFailReason = 46
	ConnectionDisconnectFailReasonInvalidTenant                                 ConnectionDisconnectFailReason = 47
	ConnectionDisconnectFailReasonUnknownPacket                                 ConnectionDisconnectFailReason = 48
	ConnectionDisconnectFailReasonUnexpectedPacket                              ConnectionDisconnectFailReason = 49
	ConnectionDisconnectFailReasonInvalidCommandRequestPacket                   ConnectionDisconnectFailReason = 50
	ConnectionDisconnectFailReasonHostSuspended                                 ConnectionDisconnectFailReason = 51
	ConnectionDisconnectFailReasonLoginPacketNoRequest                          ConnectionDisconnectFailReason = 52
	ConnectionDisconnectFailReasonLoginPacketNoCert                             ConnectionDisconnectFailReason = 53
	ConnectionDisconnectFailReasonMissingClient                                 ConnectionDisconnectFailReason = 54
	ConnectionDisconnectFailReasonKicked                                        ConnectionDisconnectFailReason = 55
	ConnectionDisconnectFailReasonKickedForExploit                              ConnectionDisconnectFailReason = 56
	ConnectionDisconnectFailReasonKickedForIdle                                 ConnectionDisconnectFailReason = 57
	ConnectionDisconnectFailReasonResourcePackProblem                           ConnectionDisconnectFailReason = 58
	ConnectionDisconnectFailReasonIncompatiblePack                              ConnectionDisconnectFailReason = 59
	ConnectionDisconnectFailReasonOutOfStorage                                  ConnectionDisconnectFailReason = 60
	ConnectionDisconnectFailReasonInvalidLevel                                  ConnectionDisconnectFailReason = 61
	ConnectionDisconnectFailReasonDisconnectPacket                              ConnectionDisconnectFailReason = 62
	ConnectionDisconnectFailReasonBlockMismatch                                 ConnectionDisconnectFailReason = 63
	ConnectionDisconnectFailReasonInvalidHeights                                ConnectionDisconnectFailReason = 64
	ConnectionDisconnectFailReasonInvalidWidths                                 ConnectionDisconnectFailReason = 65
	ConnectionDisconnectFailReasonConnectionLost                                ConnectionDisconnectFailReason = 66
	ConnectionDisconnectFailReasonZombieConnection                              ConnectionDisconnectFailReason = 67
	ConnectionDisconnectFailReasonShutdown                                      ConnectionDisconnectFailReason = 68
	ConnectionDisconnectFailReasonReasonNotSet                                  ConnectionDisconnectFailReason = 69
	ConnectionDisconnectFailReasonLoadingStateTimeout                           ConnectionDisconnectFailReason = 70
	ConnectionDisconnectFailReasonResourcePackLoadingFailed                     ConnectionDisconnectFailReason = 71
	ConnectionDisconnectFailReasonSearchingForSessionLoadingScreenFailed        ConnectionDisconnectFailReason = 72
	ConnectionDisconnectFailReasonNetherNetProtocolVersion                      ConnectionDisconnectFailReason = 73
	ConnectionDisconnectFailReasonSubsystemStatusError                          ConnectionDisconnectFailReason = 74
	ConnectionDisconnectFailReasonEmptyAuthFromDiscovery                        ConnectionDisconnectFailReason = 75
	ConnectionDisconnectFailReasonEmptyUrlFromDiscovery                         ConnectionDisconnectFailReason = 76
	ConnectionDisconnectFailReasonExpiredAuthFromDiscovery                      ConnectionDisconnectFailReason = 77
	ConnectionDisconnectFailReasonUnknownSignalServiceSignInFailure             ConnectionDisconnectFailReason = 78
	ConnectionDisconnectFailReasonXBLJoinLobbyFailure                           ConnectionDisconnectFailReason = 79
	ConnectionDisconnectFailReasonUnspecifiedClientInstanceDisconnection        ConnectionDisconnectFailReason = 80
	ConnectionDisconnectFailReasonNetherNetSessionNotFound                      ConnectionDisconnectFailReason = 81
	ConnectionDisconnectFailReasonNetherNetCreatePeerConnection                 ConnectionDisconnectFailReason = 82
	ConnectionDisconnectFailReasonNetherNetICE                                  ConnectionDisconnectFailReason = 83
	ConnectionDisconnectFailReasonNetherNetConnectRequest                       ConnectionDisconnectFailReason = 84
	ConnectionDisconnectFailReasonNetherNetConnectResponse                      ConnectionDisconnectFailReason = 85
	ConnectionDisconnectFailReasonNetherNetNegotiationTimeout                   ConnectionDisconnectFailReason = 86
	ConnectionDisconnectFailReasonNetherNetInactivityTimeout                    ConnectionDisconnectFailReason = 87
	ConnectionDisconnectFailReasonStaleConnectionBeingReplaced                  ConnectionDisconnectFailReason = 88
	ConnectionDisconnectFailReasonRealmsSessionNotFound                         ConnectionDisconnectFailReason = 89
	ConnectionDisconnectFailReasonBadPacket                                     ConnectionDisconnectFailReason = 90
	ConnectionDisconnectFailReasonNetherNetFailedToCreateOffer                  ConnectionDisconnectFailReason = 91
	ConnectionDisconnectFailReasonNetherNetFailedToCreateAnswer                 ConnectionDisconnectFailReason = 92
	ConnectionDisconnectFailReasonNetherNetFailedToSetLocalDescription          ConnectionDisconnectFailReason = 93
	ConnectionDisconnectFailReasonNetherNetFailedToSetRemoteDescription         ConnectionDisconnectFailReason = 94
	ConnectionDisconnectFailReasonNetherNetNegotiationTimeoutWaitingForResponse ConnectionDisconnectFailReason = 95
	ConnectionDisconnectFailReasonNetherNetNegotiationTimeoutWaitingForAccept   ConnectionDisconnectFailReason = 96
	ConnectionDisconnectFailReasonNetherNetIncomingConnectionIgnored            ConnectionDisconnectFailReason = 97
	ConnectionDisconnectFailReasonNetherNetSignalingParsingFailure              ConnectionDisconnectFailReason = 98
	ConnectionDisconnectFailReasonNetherNetSignalingUnknownError                ConnectionDisconnectFailReason = 99
	ConnectionDisconnectFailReasonNetherNetSignalingUnicastDeliveryFailed       ConnectionDisconnectFailReason = 100
	ConnectionDisconnectFailReasonNetherNetSignalingBroadcastDeliveryFailed     ConnectionDisconnectFailReason = 101
	ConnectionDisconnectFailReasonNetherNetSignalingGenericDeliveryFailed       ConnectionDisconnectFailReason = 102
	ConnectionDisconnectFailReasonEditorMismatchEditorWorld                     ConnectionDisconnectFailReason = 103
	ConnectionDisconnectFailReasonEditorMismatchVanillaWorld                    ConnectionDisconnectFailReason = 104
	ConnectionDisconnectFailReasonWorldTransferNotPrimaryClient                 ConnectionDisconnectFailReason = 105
	ConnectionDisconnectFailReasonINTERNALRequestServerShutdown                 ConnectionDisconnectFailReason = 106
	ConnectionDisconnectFailReasonClientGameSetupCancelled                      ConnectionDisconnectFailReason = 107
	ConnectionDisconnectFailReasonClientGameSetupFailed                         ConnectionDisconnectFailReason = 108
	ConnectionDisconnectFailReasonNoVenue                                       ConnectionDisconnectFailReason = 109
	ConnectionDisconnectFailReasonNetherNetSignalingSigninFailed                ConnectionDisconnectFailReason = 110
	ConnectionDisconnectFailReasonSessionAccessDenied                           ConnectionDisconnectFailReason = 111
	ConnectionDisconnectFailReasonServiceSigninIssue                            ConnectionDisconnectFailReason = 112
	ConnectionDisconnectFailReasonNetherNetNoSignalingChannel                   ConnectionDisconnectFailReason = 113
	ConnectionDisconnectFailReasonNetherNetNotLoggedIn                          ConnectionDisconnectFailReason = 114
	ConnectionDisconnectFailReasonNetherNetClientSignalingError                 ConnectionDisconnectFailReason = 115
	ConnectionDisconnectFailReasonSubClientLoginDisabled                        ConnectionDisconnectFailReason = 116
	ConnectionDisconnectFailReasonDeepLinkTryingToOpenDemoWorldWhileSignedIn    ConnectionDisconnectFailReason = 117
	ConnectionDisconnectFailReasonAsyncJoinTaskDenied                           ConnectionDisconnectFailReason = 118
	ConnectionDisconnectFailReasonRealmsTimelineRequired                        ConnectionDisconnectFailReason = 119
	ConnectionDisconnectFailReasonGuestWithoutHost                              ConnectionDisconnectFailReason = 120
	ConnectionDisconnectFailReasonFailedToJoinExperience                        ConnectionDisconnectFailReason = 121
	ConnectionDisconnectFailReasonNetherNetDataChannelClosed                    ConnectionDisconnectFailReason = 122
	ConnectionDisconnectFailReasonDiscoveryEnvironmentMismatch                  ConnectionDisconnectFailReason = 123
	ConnectionDisconnectFailReasonHostWithoutKeys                               ConnectionDisconnectFailReason = 124
	ConnectionDisconnectFailReasonHostSignedOut                                 ConnectionDisconnectFailReason = 125
	ConnectionDisconnectFailReasonScriptWatchdogException                       ConnectionDisconnectFailReason = 126
	ConnectionDisconnectFailReasonScriptMemoryLimitExceeded                     ConnectionDisconnectFailReason = 127
	ConnectionDisconnectFailReasonStorageLowDuringGameplay                      ConnectionDisconnectFailReason = 128
	ConnectionDisconnectFailReasonStorageFullDuringGameplay                     ConnectionDisconnectFailReason = 129
	ConnectionDisconnectFailReasonLevelStorageCorruption                        ConnectionDisconnectFailReason = 130
	ConnectionDisconnectFailReasonEditionMismatchVanillaToEdu                   ConnectionDisconnectFailReason = 131
	ConnectionDisconnectFailReasonEditionMismatchEduToVanilla                   ConnectionDisconnectFailReason = 132
	ConnectionDisconnectFailReasonEditorMismatchEditorToVanilla                 ConnectionDisconnectFailReason = 133
	ConnectionDisconnectFailReasonEditorMismatchVanillaToEditor                 ConnectionDisconnectFailReason = 134
	ConnectionDisconnectFailReasonDenyListed                                    ConnectionDisconnectFailReason = 135
	ConnectionDisconnectFailReasonNonceMissing                                  ConnectionDisconnectFailReason = 136
	ConnectionDisconnectFailReasonNonceNotFound                                 ConnectionDisconnectFailReason = 137
	ConnectionDisconnectFailReasonNonceExpired                                  ConnectionDisconnectFailReason = 138
	ConnectionDisconnectFailReasonNonceNotValid                                 ConnectionDisconnectFailReason = 139
	ConnectionDisconnectFailReasonHostDisconnected                              ConnectionDisconnectFailReason = 140
	ConnectionDisconnectFailReasonEditorJoinIntentPolicyFailure                 ConnectionDisconnectFailReason = 141
	ConnectionDisconnectFailReasonNetherNetIdentityNotAllowed                   ConnectionDisconnectFailReason = 142
	ConnectionDisconnectFailReasonInvalidName                                   ConnectionDisconnectFailReason = 143
	ConnectionDisconnectFailReasonExpiredToken                                  ConnectionDisconnectFailReason = 144
	ConnectionDisconnectFailReasonHostAcceptsNoTypeOfAuth                       ConnectionDisconnectFailReason = 145
	ConnectionDisconnectFailReasonNotAuthenticatedFastFail                      ConnectionDisconnectFailReason = 146
	ConnectionDisconnectFailReasonEditorNotAllowed                              ConnectionDisconnectFailReason = 147
)

type ContainerEnumName uint8

const (
	ContainerEnumNameAnvilInputContainer                 ContainerEnumName = 0
	ContainerEnumNameAnvilMaterialContainer              ContainerEnumName = 1
	ContainerEnumNameAnvilResultPreviewContainer         ContainerEnumName = 2
	ContainerEnumNameSmithingTableInputContainer         ContainerEnumName = 3
	ContainerEnumNameSmithingTableMaterialContainer      ContainerEnumName = 4
	ContainerEnumNameSmithingTableResultPreviewContainer ContainerEnumName = 5
	ContainerEnumNameArmorContainer                      ContainerEnumName = 6
	ContainerEnumNameLevelEntityContainer                ContainerEnumName = 7
	ContainerEnumNameBeaconPaymentContainer              ContainerEnumName = 8
	ContainerEnumNameBrewingStandInputContainer          ContainerEnumName = 9
	ContainerEnumNameBrewingStandResultContainer         ContainerEnumName = 10
	ContainerEnumNameBrewingStandFuelContainer           ContainerEnumName = 11
	ContainerEnumNameCombinedHotbarAndInventoryContainer ContainerEnumName = 12
	ContainerEnumNameCraftingInputContainer              ContainerEnumName = 13
	ContainerEnumNameCraftingOutputPreviewContainer      ContainerEnumName = 14
	ContainerEnumNameRecipeConstructionContainer         ContainerEnumName = 15
	ContainerEnumNameRecipeNatureContainer               ContainerEnumName = 16
	ContainerEnumNameRecipeItemsContainer                ContainerEnumName = 17
	ContainerEnumNameRecipeSearchContainer               ContainerEnumName = 18
	ContainerEnumNameRecipeSearchBarContainer            ContainerEnumName = 19
	ContainerEnumNameRecipeEquipmentContainer            ContainerEnumName = 20
	ContainerEnumNameRecipeBookContainer                 ContainerEnumName = 21
	ContainerEnumNameEnchantingInputContainer            ContainerEnumName = 22
	ContainerEnumNameEnchantingMaterialContainer         ContainerEnumName = 23
	ContainerEnumNameFurnaceFuelContainer                ContainerEnumName = 24
	ContainerEnumNameFurnaceIngredientContainer          ContainerEnumName = 25
	ContainerEnumNameFurnaceResultContainer              ContainerEnumName = 26
	ContainerEnumNameHorseEquipContainer                 ContainerEnumName = 27
	ContainerEnumNameHotbarContainer                     ContainerEnumName = 28
	ContainerEnumNameInventoryContainer                  ContainerEnumName = 29
	ContainerEnumNameShulkerBoxContainer                 ContainerEnumName = 30
	ContainerEnumNameTradeIngredient1Container           ContainerEnumName = 31
	ContainerEnumNameTradeIngredient2Container           ContainerEnumName = 32
	ContainerEnumNameTradeResultPreviewContainer         ContainerEnumName = 33
	ContainerEnumNameOffhandContainer                    ContainerEnumName = 34
	ContainerEnumNameCompoundCreatorInput                ContainerEnumName = 35
	ContainerEnumNameCompoundCreatorOutputPreview        ContainerEnumName = 36
	ContainerEnumNameElementConstructorOutputPreview     ContainerEnumName = 37
	ContainerEnumNameMaterialReducerInput                ContainerEnumName = 38
	ContainerEnumNameMaterialReducerOutput               ContainerEnumName = 39
	ContainerEnumNameLabTableInput                       ContainerEnumName = 40
	ContainerEnumNameLoomInputContainer                  ContainerEnumName = 41
	ContainerEnumNameLoomDyeContainer                    ContainerEnumName = 42
	ContainerEnumNameLoomMaterialContainer               ContainerEnumName = 43
	ContainerEnumNameLoomResultPreviewContainer          ContainerEnumName = 44
	ContainerEnumNameBlastFurnaceIngredientContainer     ContainerEnumName = 45
	ContainerEnumNameSmokerIngredientContainer           ContainerEnumName = 46
	ContainerEnumNameTrade2Ingredient1Container          ContainerEnumName = 47
	ContainerEnumNameTrade2Ingredient2Container          ContainerEnumName = 48
	ContainerEnumNameTrade2ResultPreviewContainer        ContainerEnumName = 49
	ContainerEnumNameGrindstoneInputContainer            ContainerEnumName = 50
	ContainerEnumNameGrindstoneAdditionalContainer       ContainerEnumName = 51
	ContainerEnumNameGrindstoneResultPreviewContainer    ContainerEnumName = 52
	ContainerEnumNameStonecutterInputContainer           ContainerEnumName = 53
	ContainerEnumNameStonecutterResultPreviewContainer   ContainerEnumName = 54
	ContainerEnumNameCartographyInputContainer           ContainerEnumName = 55
	ContainerEnumNameCartographyAdditionalContainer      ContainerEnumName = 56
	ContainerEnumNameCartographyResultPreviewContainer   ContainerEnumName = 57
	ContainerEnumNameBarrelContainer                     ContainerEnumName = 58
	ContainerEnumNameCursorContainer                     ContainerEnumName = 59
	ContainerEnumNameCreatedOutputContainer              ContainerEnumName = 60
	ContainerEnumNameSmithingTableTemplateContainer      ContainerEnumName = 61
	ContainerEnumNameCrafterLevelEntityContainer         ContainerEnumName = 62
	ContainerEnumNameDynamicContainer                    ContainerEnumName = 63
	ContainerEnumNameRecipeFoodContainer                 ContainerEnumName = 64
	ContainerEnumNameRecipeBlocksContainer               ContainerEnumName = 65
	ContainerEnumNameRecipeFurnaceItemsContainer         ContainerEnumName = 66
)

type ControlSchemeScheme uint8

const (
	ControlSchemeSchemeLockedPlayerRelativeStrafe ControlSchemeScheme = 0
	ControlSchemeSchemeCameraRelative             ControlSchemeScheme = 1
	ControlSchemeSchemeCameraRelativeStrafe       ControlSchemeScheme = 2
	ControlSchemeSchemePlayerRelative             ControlSchemeScheme = 3
	ControlSchemeSchemePlayerRelativeStrafe       ControlSchemeScheme = 4
)

type CoordinateEvaluationOrder int32

const (
	CoordinateEvaluationOrderXYZ CoordinateEvaluationOrder = 0
	CoordinateEvaluationOrderXZY CoordinateEvaluationOrder = 1
	CoordinateEvaluationOrderYXZ CoordinateEvaluationOrder = 2
	CoordinateEvaluationOrderYZX CoordinateEvaluationOrder = 3
	CoordinateEvaluationOrderZXY CoordinateEvaluationOrder = 4
	CoordinateEvaluationOrderZYX CoordinateEvaluationOrder = 5
)

type CreativeItemCategory uint8

const (
	CreativeItemCategoryConstruction    CreativeItemCategory = 1
	CreativeItemCategoryNature          CreativeItemCategory = 2
	CreativeItemCategoryEquipment       CreativeItemCategory = 3
	CreativeItemCategoryItems           CreativeItemCategory = 4
	CreativeItemCategoryItemCommandOnly CreativeItemCategory = 5
)

type DataItemType uint8

const (
	DataItemTypeByte        DataItemType = 0
	DataItemTypeShort       DataItemType = 1
	DataItemTypeInt         DataItemType = 2
	DataItemTypeFloat       DataItemType = 3
	DataItemTypeString      DataItemType = 4
	DataItemTypeCompoundTag DataItemType = 5
	DataItemTypePos         DataItemType = 6
	DataItemTypeInt64       DataItemType = 7
	DataItemTypeVec3        DataItemType = 8
)

type EditorWorldType int32

const (
	EditorWorldTypeNonEditor          EditorWorldType = 0
	EditorWorldTypeEditorProject      EditorWorldType = 1
	EditorWorldTypeEditorTestLevel    EditorWorldType = 2
	EditorWorldTypeEditorRealmsUpload EditorWorldType = 3
)

type EducationEditionOffer uint32

const (
	EducationEditionOfferNone            EducationEditionOffer = 0
	EducationEditionOfferRestOfWorld     EducationEditionOffer = 1
	EducationEditionOfferChinaDeprecated EducationEditionOffer = 2
)

type EnchantType uint8

const (
	EnchantTypeProtection           EnchantType = 0
	EnchantTypeFireProtection       EnchantType = 1
	EnchantTypeFeatherFalling       EnchantType = 2
	EnchantTypeBlastProtection      EnchantType = 3
	EnchantTypeProjectileProtection EnchantType = 4
	EnchantTypeThorns               EnchantType = 5
	EnchantTypeRespiration          EnchantType = 6
	EnchantTypeDepthStrider         EnchantType = 7
	EnchantTypeAquaAffinity         EnchantType = 8
	EnchantTypeSharpness            EnchantType = 9
	EnchantTypeSmite                EnchantType = 10
	EnchantTypeBaneOfArthropods     EnchantType = 11
	EnchantTypeKnockback            EnchantType = 12
	EnchantTypeFireAspect           EnchantType = 13
	EnchantTypeLooting              EnchantType = 14
	EnchantTypeEfficiency           EnchantType = 15
	EnchantTypeSilkTouch            EnchantType = 16
	EnchantTypeUnbreaking           EnchantType = 17
	EnchantTypeFortune              EnchantType = 18
	EnchantTypePower                EnchantType = 19
	EnchantTypePunch                EnchantType = 20
	EnchantTypeFlame                EnchantType = 21
	EnchantTypeInfinity             EnchantType = 22
	EnchantTypeLuckOfTheSea         EnchantType = 23
	EnchantTypeLure                 EnchantType = 24
	EnchantTypeFrostWalker          EnchantType = 25
	EnchantTypeMending              EnchantType = 26
	EnchantTypeCurseOfBinding       EnchantType = 27
	EnchantTypeCurseOfVanishing     EnchantType = 28
	EnchantTypeImpaling             EnchantType = 29
	EnchantTypeRiptide              EnchantType = 30
	EnchantTypeLoyalty              EnchantType = 31
	EnchantTypeChanneling           EnchantType = 32
	EnchantTypeMultishot            EnchantType = 33
	EnchantTypePiercing             EnchantType = 34
	EnchantTypeQuickCharge          EnchantType = 35
	EnchantTypeSoulSpeed            EnchantType = 36
	EnchantTypeSwiftSneak           EnchantType = 37
	EnchantTypeWindBurst            EnchantType = 38
	EnchantTypeDensity              EnchantType = 39
	EnchantTypeBreach               EnchantType = 40
	EnchantTypeLunge                EnchantType = 41
	EnchantTypeNumEnchantments      EnchantType = 42
	EnchantTypeInvalidEnchantment   EnchantType = 43
)

type GameType int32

const (
	GameTypeUndefined GameType = -1
	GameTypeSurvival  GameType = 0
	GameTypeCreative  GameType = 1
	GameTypeAdventure GameType = 2
	GameTypeDefault   GameType = 5
	GameTypeSpectator GameType = 6
)

type GeneratorType int32

const (
	GeneratorTypeLegacy    GeneratorType = 0
	GeneratorTypeOverworld GeneratorType = 1
	GeneratorTypeFlat      GeneratorType = 2
	GeneratorTypeNether    GeneratorType = 3
	GeneratorTypeTheEnd    GeneratorType = 4
	GeneratorTypeVoid      GeneratorType = 5
	GeneratorTypeUndefined GeneratorType = 6
)

type GraphicsMode uint8

const (
	GraphicsModeSimple    GraphicsMode = 0
	GraphicsModeFancy     GraphicsMode = 1
	GraphicsModeAdvanced  GraphicsMode = 2
	GraphicsModeRayTraced GraphicsMode = 3
)

type GraphicsOverrideParameterType uint8

const (
	GraphicsOverrideParameterTypeSkyZenithColor          GraphicsOverrideParameterType = 0
	GraphicsOverrideParameterTypeSkyHorizonColor         GraphicsOverrideParameterType = 1
	GraphicsOverrideParameterTypeHorizonBlendMin         GraphicsOverrideParameterType = 2
	GraphicsOverrideParameterTypeHorizonBlendMax         GraphicsOverrideParameterType = 3
	GraphicsOverrideParameterTypeHorizonBlendStart       GraphicsOverrideParameterType = 4
	GraphicsOverrideParameterTypeHorizonBlendMieStart    GraphicsOverrideParameterType = 5
	GraphicsOverrideParameterTypeRayleighStrength        GraphicsOverrideParameterType = 6
	GraphicsOverrideParameterTypeSunMieStrength          GraphicsOverrideParameterType = 7
	GraphicsOverrideParameterTypeMoonMieStrength         GraphicsOverrideParameterType = 8
	GraphicsOverrideParameterTypeSunGlareShape           GraphicsOverrideParameterType = 9
	GraphicsOverrideParameterTypeChlorophyll             GraphicsOverrideParameterType = 10
	GraphicsOverrideParameterTypeCDOM                    GraphicsOverrideParameterType = 11
	GraphicsOverrideParameterTypeSuspendedSediment       GraphicsOverrideParameterType = 12
	GraphicsOverrideParameterTypeWavesDepth              GraphicsOverrideParameterType = 13
	GraphicsOverrideParameterTypeWavesFrequency          GraphicsOverrideParameterType = 14
	GraphicsOverrideParameterTypeWavesFrequencyScaling   GraphicsOverrideParameterType = 15
	GraphicsOverrideParameterTypeWavesSpeed              GraphicsOverrideParameterType = 16
	GraphicsOverrideParameterTypeWavesSpeedScaling       GraphicsOverrideParameterType = 17
	GraphicsOverrideParameterTypeWavesShape              GraphicsOverrideParameterType = 18
	GraphicsOverrideParameterTypeWavesOctaves            GraphicsOverrideParameterType = 19
	GraphicsOverrideParameterTypeWavesMix                GraphicsOverrideParameterType = 20
	GraphicsOverrideParameterTypeWavesPull               GraphicsOverrideParameterType = 21
	GraphicsOverrideParameterTypeWavesDirectionIncrement GraphicsOverrideParameterType = 22
	GraphicsOverrideParameterTypeMidtonesContrast        GraphicsOverrideParameterType = 23
	GraphicsOverrideParameterTypeHighlightsContrast      GraphicsOverrideParameterType = 24
	GraphicsOverrideParameterTypeShadowsContrast         GraphicsOverrideParameterType = 25
	GraphicsOverrideParameterTypeHighlightsGain          GraphicsOverrideParameterType = 26
	GraphicsOverrideParameterTypeHighlightsGamma         GraphicsOverrideParameterType = 27
	GraphicsOverrideParameterTypeHighlightsOffset        GraphicsOverrideParameterType = 28
	GraphicsOverrideParameterTypeHighlightsSaturation    GraphicsOverrideParameterType = 29
	GraphicsOverrideParameterTypeMidtonesGain            GraphicsOverrideParameterType = 30
	GraphicsOverrideParameterTypeMidtonesGamma           GraphicsOverrideParameterType = 31
	GraphicsOverrideParameterTypeMidtonesOffset          GraphicsOverrideParameterType = 32
	GraphicsOverrideParameterTypeMidtonesSaturation      GraphicsOverrideParameterType = 33
	GraphicsOverrideParameterTypeShadowsGain             GraphicsOverrideParameterType = 34
	GraphicsOverrideParameterTypeShadowsGamma            GraphicsOverrideParameterType = 35
	GraphicsOverrideParameterTypeShadowsOffset           GraphicsOverrideParameterType = 36
	GraphicsOverrideParameterTypeShadowsSaturation       GraphicsOverrideParameterType = 37
	GraphicsOverrideParameterTypeHighlightsMin           GraphicsOverrideParameterType = 38
	GraphicsOverrideParameterTypeShadowsMax              GraphicsOverrideParameterType = 39
	GraphicsOverrideParameterTypeTemperature             GraphicsOverrideParameterType = 40
	GraphicsOverrideParameterTypeSunColor                GraphicsOverrideParameterType = 41
	GraphicsOverrideParameterTypeSunIlluminance          GraphicsOverrideParameterType = 42
	GraphicsOverrideParameterTypeMoonColor               GraphicsOverrideParameterType = 43
	GraphicsOverrideParameterTypeMoonIlluminance         GraphicsOverrideParameterType = 44
	GraphicsOverrideParameterTypeFlashColor              GraphicsOverrideParameterType = 45
	GraphicsOverrideParameterTypeFlashIlluminance        GraphicsOverrideParameterType = 46
	GraphicsOverrideParameterTypeAmbientColor            GraphicsOverrideParameterType = 47
	GraphicsOverrideParameterTypeAmbientIlluminance      GraphicsOverrideParameterType = 48
	GraphicsOverrideParameterTypeEmissiveDesaturation    GraphicsOverrideParameterType = 49
	GraphicsOverrideParameterTypeSkyIntensity            GraphicsOverrideParameterType = 50
	GraphicsOverrideParameterTypeOrbitalOffsetDegrees    GraphicsOverrideParameterType = 51
)

type HudElement int32

const (
	HudElementPaperDoll     HudElement = 0
	HudElementArmor         HudElement = 1
	HudElementToolTips      HudElement = 2
	HudElementTouchControls HudElement = 3
	HudElementCrosshair     HudElement = 4
	HudElementHotBar        HudElement = 5
	HudElementHealth        HudElement = 6
	HudElementProgressBar   HudElement = 7
	HudElementHunger        HudElement = 8
	HudElementAirBubbles    HudElement = 9
	HudElementHorseHealth   HudElement = 10
	HudElementStatusEffects HudElement = 11
	HudElementItemText      HudElement = 12
)

type HudVisibility int32

const (
	HudVisibilityHide  HudVisibility = 0
	HudVisibilityReset HudVisibility = 1
)

type InputMode uint32

const (
	InputModeUndefined        InputMode = 0
	InputModeMouse            InputMode = 1
	InputModeTouch            InputMode = 2
	InputModeGamePad          InputMode = 3
	InputModeMotionController InputMode = 4
	InputModeCount            InputMode = 5
)

type InteractAction uint8

const (
	InteractActionInvalid        InteractAction = 0
	InteractActionStopRiding     InteractAction = 3
	InteractActionInteractUpdate InteractAction = 4
	InteractActionNpcOpen        InteractAction = 5
	InteractActionOpenInventory  InteractAction = 6
)

type InventoryLayout int32

const (
	InventoryLayoutNone           InventoryLayout = 0
	InventoryLayoutInventoryOnly  InventoryLayout = 1
	InventoryLayoutDefault        InventoryLayout = 2
	InventoryLayoutRecipeBookOnly InventoryLayout = 3
)

type InventoryLeftTabIndex int32

const (
	InventoryLeftTabIndexNone               InventoryLeftTabIndex = 0
	InventoryLeftTabIndexRecipeConstruction InventoryLeftTabIndex = 1
	InventoryLeftTabIndexRecipeEquipment    InventoryLeftTabIndex = 2
	InventoryLeftTabIndexRecipeItems        InventoryLeftTabIndex = 3
	InventoryLeftTabIndexRecipeNature       InventoryLeftTabIndex = 4
	InventoryLeftTabIndexRecipeSearch       InventoryLeftTabIndex = 5
	InventoryLeftTabIndexSurvival           InventoryLeftTabIndex = 6
)

type InventoryRightTabIndex int32

const (
	InventoryRightTabIndexNone       InventoryRightTabIndex = 0
	InventoryRightTabIndexFullScreen InventoryRightTabIndex = 1
	InventoryRightTabIndexCrafting   InventoryRightTabIndex = 2
	InventoryRightTabIndexArmor      InventoryRightTabIndex = 3
)

type InventorySourceInventorySourceFlags uint32

const (
	InventorySourceInventorySourceFlagsNoFlag                 InventorySourceInventorySourceFlags = 0
	InventorySourceInventorySourceFlagsWorldInteractionRandom InventorySourceInventorySourceFlags = 1
)

type InventorySourceType uint32

const (
	InventorySourceTypeContainerInventory        InventorySourceType = 0
	InventorySourceTypeGlobalInventory           InventorySourceType = 1
	InventorySourceTypeWorldInteraction          InventorySourceType = 2
	InventorySourceTypeCreativeInventory         InventorySourceType = 3
	InventorySourceTypeNonImplementedFeatureTODO InventorySourceType = 99999
)

type ItemReleaseInventoryTransactionActionType int32

const (
	ItemReleaseInventoryTransactionActionTypeRelease ItemReleaseInventoryTransactionActionType = 0
	ItemReleaseInventoryTransactionActionTypeUse     ItemReleaseInventoryTransactionActionType = 1
)

type ItemStackNetResult uint8

const (
	ItemStackNetResultSuccess                                          ItemStackNetResult = 0
	ItemStackNetResultError                                            ItemStackNetResult = 1
	ItemStackNetResultInvalidRequestActionType                         ItemStackNetResult = 2
	ItemStackNetResultActionRequestNotAllowed                          ItemStackNetResult = 3
	ItemStackNetResultScreenHandlerEndRequestFailed                    ItemStackNetResult = 4
	ItemStackNetResultItemRequestActionHandlerCommitFailed             ItemStackNetResult = 5
	ItemStackNetResultInvalidRequestCraftActionType                    ItemStackNetResult = 6
	ItemStackNetResultInvalidCraftRequest                              ItemStackNetResult = 7
	ItemStackNetResultInvalidCraftRequestScreen                        ItemStackNetResult = 8
	ItemStackNetResultInvalidCraftResult                               ItemStackNetResult = 9
	ItemStackNetResultInvalidCraftResultIndex                          ItemStackNetResult = 10
	ItemStackNetResultInvalidCraftResultItem                           ItemStackNetResult = 11
	ItemStackNetResultInvalidItemNetId                                 ItemStackNetResult = 12
	ItemStackNetResultMissingCreatedOutputContainer                    ItemStackNetResult = 13
	ItemStackNetResultFailedToSetCreatedItemOutputSlot                 ItemStackNetResult = 14
	ItemStackNetResultRequestAlreadyInProgress                         ItemStackNetResult = 15
	ItemStackNetResultFailedToInitSparseContainer                      ItemStackNetResult = 16
	ItemStackNetResultResultTransferFailed                             ItemStackNetResult = 17
	ItemStackNetResultExpectedItemSlotNotFullyConsumed                 ItemStackNetResult = 18
	ItemStackNetResultExpectedAnywhereItemNotFullyConsumed             ItemStackNetResult = 19
	ItemStackNetResultItemAlreadyConsumedFromSlot                      ItemStackNetResult = 20
	ItemStackNetResultConsumedTooMuchFromSlot                          ItemStackNetResult = 21
	ItemStackNetResultMismatchSlotExpectedConsumedItem                 ItemStackNetResult = 22
	ItemStackNetResultMismatchSlotExpectedConsumedItemNetIdVariant     ItemStackNetResult = 23
	ItemStackNetResultFailedToMatchExpectedSlotConsumedItem            ItemStackNetResult = 24
	ItemStackNetResultFailedToMatchExpectedAllowedAnywhereConsumedItem ItemStackNetResult = 25
	ItemStackNetResultConsumedItemOutOfAllowedSlotRange                ItemStackNetResult = 26
	ItemStackNetResultConsumedItemNotAllowed                           ItemStackNetResult = 27
	ItemStackNetResultPlayerNotInCreativeMode                          ItemStackNetResult = 28
	ItemStackNetResultInvalidExperimentalRecipeRequest                 ItemStackNetResult = 29
	ItemStackNetResultFailedToCraftCreative                            ItemStackNetResult = 30
	ItemStackNetResultFailedToGetLevelRecipe                           ItemStackNetResult = 31
	ItemStackNetResultFailedToFindRecipeByNetId                        ItemStackNetResult = 32
	ItemStackNetResultMismatchedCraftingSize                           ItemStackNetResult = 33
	ItemStackNetResultMissingInputSparseContainer                      ItemStackNetResult = 34
	ItemStackNetResultMismatchedRecipeForInputGridItems                ItemStackNetResult = 35
	ItemStackNetResultEmptyCraftResults                                ItemStackNetResult = 36
	ItemStackNetResultFailedToEnchant                                  ItemStackNetResult = 37
	ItemStackNetResultMissingInputItem                                 ItemStackNetResult = 38
	ItemStackNetResultInsufficientPlayerLevelToEnchant                 ItemStackNetResult = 39
	ItemStackNetResultMissingMaterialItem                              ItemStackNetResult = 40
	ItemStackNetResultMissingActor                                     ItemStackNetResult = 41
	ItemStackNetResultUnknownPrimaryEffect                             ItemStackNetResult = 42
	ItemStackNetResultPrimaryEffectOutOfRange                          ItemStackNetResult = 43
	ItemStackNetResultPrimaryEffectUnavailable                         ItemStackNetResult = 44
	ItemStackNetResultSecondaryEffectOutOfRange                        ItemStackNetResult = 45
	ItemStackNetResultSecondaryEffectUnavailable                       ItemStackNetResult = 46
	ItemStackNetResultDstContainerEqualToCreatedOutputContainer        ItemStackNetResult = 47
	ItemStackNetResultDstContainerAndSlotEqualToSrcContainerAndSlot    ItemStackNetResult = 48
	ItemStackNetResultFailedToValidateSrcSlot                          ItemStackNetResult = 49
	ItemStackNetResultFailedToValidateDstSlot                          ItemStackNetResult = 50
	ItemStackNetResultInvalidAdjustedAmount                            ItemStackNetResult = 51
	ItemStackNetResultInvalidItemSetType                               ItemStackNetResult = 52
	ItemStackNetResultInvalidTransferAmount                            ItemStackNetResult = 53
	ItemStackNetResultCannotSwapItem                                   ItemStackNetResult = 54
	ItemStackNetResultCannotPlaceItem                                  ItemStackNetResult = 55
	ItemStackNetResultUnhandledItemSetType                             ItemStackNetResult = 56
	ItemStackNetResultInvalidRemovedAmount                             ItemStackNetResult = 57
	ItemStackNetResultInvalidRegion                                    ItemStackNetResult = 58
	ItemStackNetResultCannotDropItem                                   ItemStackNetResult = 59
	ItemStackNetResultCannotDestroyItem                                ItemStackNetResult = 60
	ItemStackNetResultInvalidSourceContainer                           ItemStackNetResult = 61
	ItemStackNetResultItemNotConsumed                                  ItemStackNetResult = 62
	ItemStackNetResultInvalidNumCrafts                                 ItemStackNetResult = 63
	ItemStackNetResultInvalidCraftResultStackSize                      ItemStackNetResult = 64
	ItemStackNetResultCannotRemoveItem                                 ItemStackNetResult = 65
	ItemStackNetResultCannotConsumeItem                                ItemStackNetResult = 66
	ItemStackNetResultScreenStackError                                 ItemStackNetResult = 67
)

type ItemStackRequestActionType uint8

const (
	ItemStackRequestActionTypeTake                     ItemStackRequestActionType = 0
	ItemStackRequestActionTypePlace                    ItemStackRequestActionType = 1
	ItemStackRequestActionTypeSwap                     ItemStackRequestActionType = 2
	ItemStackRequestActionTypeDrop                     ItemStackRequestActionType = 3
	ItemStackRequestActionTypeDestroy                  ItemStackRequestActionType = 4
	ItemStackRequestActionTypeConsume                  ItemStackRequestActionType = 5
	ItemStackRequestActionTypeCreate                   ItemStackRequestActionType = 6
	ItemStackRequestActionTypePlaceInItemContainer     ItemStackRequestActionType = 7
	ItemStackRequestActionTypeTakeFromItemContainer    ItemStackRequestActionType = 8
	ItemStackRequestActionTypeScreenLabTableCombine    ItemStackRequestActionType = 9
	ItemStackRequestActionTypeScreenBeaconPayment      ItemStackRequestActionType = 10
	ItemStackRequestActionTypeScreenHUDMineBlock       ItemStackRequestActionType = 11
	ItemStackRequestActionTypeCraftRecipe              ItemStackRequestActionType = 12
	ItemStackRequestActionTypeCraftRecipeAuto          ItemStackRequestActionType = 13
	ItemStackRequestActionTypeCraftCreative            ItemStackRequestActionType = 14
	ItemStackRequestActionTypeCraftRecipeOptional      ItemStackRequestActionType = 15
	ItemStackRequestActionTypeCraftRepairAndDisenchant ItemStackRequestActionType = 16
	ItemStackRequestActionTypeCraftLoom                ItemStackRequestActionType = 17
	ItemStackRequestActionTypeCraftNonImplemented      ItemStackRequestActionType = 18
	ItemStackRequestActionTypeCraftResults             ItemStackRequestActionType = 19
)

type ItemStackRequestCerealItemDescriptorType uint8

const (
	ItemStackRequestCerealItemDescriptorTypeEmpty    ItemStackRequestCerealItemDescriptorType = 0
	ItemStackRequestCerealItemDescriptorTypeItemName ItemStackRequestCerealItemDescriptorType = 1
	ItemStackRequestCerealItemDescriptorTypeMolang   ItemStackRequestCerealItemDescriptorType = 2
	ItemStackRequestCerealItemDescriptorTypeItemTag  ItemStackRequestCerealItemDescriptorType = 3
)

type ItemUseInventoryTransactionActionType int32

const (
	ItemUseInventoryTransactionActionTypePlace       ItemUseInventoryTransactionActionType = 0
	ItemUseInventoryTransactionActionTypeUse         ItemUseInventoryTransactionActionType = 1
	ItemUseInventoryTransactionActionTypeDestroy     ItemUseInventoryTransactionActionType = 2
	ItemUseInventoryTransactionActionTypeUseAsAttack ItemUseInventoryTransactionActionType = 3
)

type ItemUseInventoryTransactionClientCooldownState uint8

const (
	ItemUseInventoryTransactionClientCooldownStateOff ItemUseInventoryTransactionClientCooldownState = 0
	ItemUseInventoryTransactionClientCooldownStateOn  ItemUseInventoryTransactionClientCooldownState = 1
)

type ItemUseInventoryTransactionPredictedResult uint8

const (
	ItemUseInventoryTransactionPredictedResultFailure ItemUseInventoryTransactionPredictedResult = 0
	ItemUseInventoryTransactionPredictedResultSuccess ItemUseInventoryTransactionPredictedResult = 1
)

type ItemUseInventoryTransactionTriggerType uint8

const (
	ItemUseInventoryTransactionTriggerTypeUnknown        ItemUseInventoryTransactionTriggerType = 0
	ItemUseInventoryTransactionTriggerTypePlayerInput    ItemUseInventoryTransactionTriggerType = 1
	ItemUseInventoryTransactionTriggerTypeSimulationTick ItemUseInventoryTransactionTriggerType = 2
)

type ItemUseOnActorInventoryTransactionActionType int32

const (
	ItemUseOnActorInventoryTransactionActionTypeInteract     ItemUseOnActorInventoryTransactionActionType = 0
	ItemUseOnActorInventoryTransactionActionTypeAttack       ItemUseOnActorInventoryTransactionActionType = 1
	ItemUseOnActorInventoryTransactionActionTypeItemInteract ItemUseOnActorInventoryTransactionActionType = 2
)

type ItemVersion int32

const (
	ItemVersionLegacy     ItemVersion = 0
	ItemVersionDataDriven ItemVersion = 1
	ItemVersionNone       ItemVersion = 2
)

type LabTableReactionType uint8

const (
	LabTableReactionTypeNone               LabTableReactionType = 0
	LabTableReactionTypeIceBomb            LabTableReactionType = 1
	LabTableReactionTypeBleach             LabTableReactionType = 2
	LabTableReactionTypeElephantToothpaste LabTableReactionType = 3
	LabTableReactionTypeFertilizer         LabTableReactionType = 4
	LabTableReactionTypeHeatBlock          LabTableReactionType = 5
	LabTableReactionTypeMagnesiumSalts     LabTableReactionType = 6
	LabTableReactionTypeMiscFire           LabTableReactionType = 7
	LabTableReactionTypeMiscExplosion      LabTableReactionType = 8
	LabTableReactionTypeMiscLava           LabTableReactionType = 9
	LabTableReactionTypeMiscMystical       LabTableReactionType = 10
	LabTableReactionTypeMiscSmoke          LabTableReactionType = 11
	LabTableReactionTypeMiscLargeSmoke     LabTableReactionType = 12
)

type LabTableType uint8

const (
	LabTableTypeStartCombine  LabTableType = 0
	LabTableTypeStartReaction LabTableType = 1
	LabTableTypeReset         LabTableType = 2
)

type LegacyArmorSlot int32

const (
	LegacyArmorSlotHead  LegacyArmorSlot = 0
	LegacyArmorSlotTorso LegacyArmorSlot = 1
	LegacyArmorSlotLegs  LegacyArmorSlot = 2
	LegacyArmorSlotFeet  LegacyArmorSlot = 3
	LegacyArmorSlotBody  LegacyArmorSlot = 4
)

type LegacyDifficulty int32

const (
	LegacyDifficultyPeaceful LegacyDifficulty = 0
	LegacyDifficultyEasy     LegacyDifficulty = 1
	LegacyDifficultyNormal   LegacyDifficulty = 2
	LegacyDifficultyHard     LegacyDifficulty = 3
	LegacyDifficultyCount    LegacyDifficulty = 4
	LegacyDifficultyUnknown  LegacyDifficulty = 5
)

type LegacyTelemetryEventType int32

const (
	LegacyTelemetryEventTypeAchievement                     LegacyTelemetryEventType = 0
	LegacyTelemetryEventTypeInteraction                     LegacyTelemetryEventType = 1
	LegacyTelemetryEventTypePortalCreated                   LegacyTelemetryEventType = 2
	LegacyTelemetryEventTypePortalUsed                      LegacyTelemetryEventType = 3
	LegacyTelemetryEventTypeMobKilled                       LegacyTelemetryEventType = 4
	LegacyTelemetryEventTypeCauldronUsed                    LegacyTelemetryEventType = 5
	LegacyTelemetryEventTypePlayerDied                      LegacyTelemetryEventType = 6
	LegacyTelemetryEventTypeBossKilled                      LegacyTelemetryEventType = 7
	LegacyTelemetryEventTypeAgentCommandOBSOLETE            LegacyTelemetryEventType = 8
	LegacyTelemetryEventTypeAgentCreated                    LegacyTelemetryEventType = 9
	LegacyTelemetryEventTypePatternRemovedOBSOLETE          LegacyTelemetryEventType = 10
	LegacyTelemetryEventTypeSlashCommand                    LegacyTelemetryEventType = 11
	LegacyTelemetryEventTypeFishBucketedOBSOLETE            LegacyTelemetryEventType = 12
	LegacyTelemetryEventTypeMobBorn                         LegacyTelemetryEventType = 13
	LegacyTelemetryEventTypePetDiedOBSOLETE                 LegacyTelemetryEventType = 14
	LegacyTelemetryEventTypePOICauldronUsed                 LegacyTelemetryEventType = 15
	LegacyTelemetryEventTypeComposterUsed                   LegacyTelemetryEventType = 16
	LegacyTelemetryEventTypeBellUsed                        LegacyTelemetryEventType = 17
	LegacyTelemetryEventTypeActorDefinition                 LegacyTelemetryEventType = 18
	LegacyTelemetryEventTypeRaidUpdate                      LegacyTelemetryEventType = 19
	LegacyTelemetryEventTypePlayerMovementAnomalyOBSOLETE   LegacyTelemetryEventType = 20
	LegacyTelemetryEventTypePlayerMovementCorrectedOBSOLETE LegacyTelemetryEventType = 21
	LegacyTelemetryEventTypeHoneyHarvested                  LegacyTelemetryEventType = 22
	LegacyTelemetryEventTypeTargetBlockHit                  LegacyTelemetryEventType = 23
	LegacyTelemetryEventTypePiglinBarter                    LegacyTelemetryEventType = 24
	LegacyTelemetryEventTypePlayerWaxedOrUnwaxedCopper      LegacyTelemetryEventType = 25
	LegacyTelemetryEventTypeCodeBuilderRuntimeAction        LegacyTelemetryEventType = 26
	LegacyTelemetryEventTypeCodeBuilderScoreboard           LegacyTelemetryEventType = 27
	LegacyTelemetryEventTypeStriderRiddenInLavaInOverworld  LegacyTelemetryEventType = 28
	LegacyTelemetryEventTypeSneakCloseToSculkSensor         LegacyTelemetryEventType = 29
	LegacyTelemetryEventTypeCarefulRestoration              LegacyTelemetryEventType = 30
	LegacyTelemetryEventTypeItemUsed                        LegacyTelemetryEventType = 31
)

type MapDecorationType int8

const (
	MapDecorationTypeMarkerWhite      MapDecorationType = 0
	MapDecorationTypeMarkerGreen      MapDecorationType = 1
	MapDecorationTypeMarkerRed        MapDecorationType = 2
	MapDecorationTypeMarkerBlue       MapDecorationType = 3
	MapDecorationTypeXWhite           MapDecorationType = 4
	MapDecorationTypeTriangleRed      MapDecorationType = 5
	MapDecorationTypeSquareWhite      MapDecorationType = 6
	MapDecorationTypeMarkerSign       MapDecorationType = 7
	MapDecorationTypeMarkerPink       MapDecorationType = 8
	MapDecorationTypeMarkerOrange     MapDecorationType = 9
	MapDecorationTypeMarkerYellow     MapDecorationType = 10
	MapDecorationTypeMarkerTeal       MapDecorationType = 11
	MapDecorationTypeTriangleGreen    MapDecorationType = 12
	MapDecorationTypeSmallSquareWhite MapDecorationType = 13
	MapDecorationTypeMansion          MapDecorationType = 14
	MapDecorationTypeMonument         MapDecorationType = 15
	MapDecorationTypeNoDraw           MapDecorationType = 16
	MapDecorationTypeVillageDesert    MapDecorationType = 17
	MapDecorationTypeVillagePlains    MapDecorationType = 18
	MapDecorationTypeVillageSavanna   MapDecorationType = 19
	MapDecorationTypeVillageSnowy     MapDecorationType = 20
	MapDecorationTypeVillageTaiga     MapDecorationType = 21
	MapDecorationTypeJungleTemple     MapDecorationType = 22
	MapDecorationTypeWitchHut         MapDecorationType = 23
	MapDecorationTypeTrialChambers    MapDecorationType = 24
	MapDecorationTypeCount            MapDecorationType = 25
)

type MapItemTrackedActorType int32

const (
	MapItemTrackedActorTypeEntity      MapItemTrackedActorType = 0
	MapItemTrackedActorTypeBlockEntity MapItemTrackedActorType = 1
	MapItemTrackedActorTypeOther       MapItemTrackedActorType = 2
)

type MemoryMemoryCategory uint8

const (
	MemoryMemoryCategoryUnknown                               MemoryMemoryCategory = 0
	MemoryMemoryCategoryInvalidSizeUnknown                    MemoryMemoryCategory = 1
	MemoryMemoryCategoryActor                                 MemoryMemoryCategory = 2
	MemoryMemoryCategoryActorAnimation                        MemoryMemoryCategory = 3
	MemoryMemoryCategoryActorRendering                        MemoryMemoryCategory = 4
	MemoryMemoryCategoryBlockTickingQueues                    MemoryMemoryCategory = 5
	MemoryMemoryCategoryBiomeStorage                          MemoryMemoryCategory = 6
	MemoryMemoryCategoryBlobs                                 MemoryMemoryCategory = 7
	MemoryMemoryCategoryCereal                                MemoryMemoryCategory = 8
	MemoryMemoryCategoryCircuitSystem                         MemoryMemoryCategory = 9
	MemoryMemoryCategoryClient                                MemoryMemoryCategory = 10
	MemoryMemoryCategoryCommands                              MemoryMemoryCategory = 11
	MemoryMemoryCategoryDBStorage                             MemoryMemoryCategory = 12
	MemoryMemoryCategoryDebug                                 MemoryMemoryCategory = 13
	MemoryMemoryCategoryDocumentation                         MemoryMemoryCategory = 14
	MemoryMemoryCategoryECSSystems                            MemoryMemoryCategory = 15
	MemoryMemoryCategoryFMOD                                  MemoryMemoryCategory = 16
	MemoryMemoryCategoryFonts                                 MemoryMemoryCategory = 17
	MemoryMemoryCategoryImGui                                 MemoryMemoryCategory = 18
	MemoryMemoryCategoryInput                                 MemoryMemoryCategory = 19
	MemoryMemoryCategoryJsonUI                                MemoryMemoryCategory = 20
	MemoryMemoryCategoryJsonUIControlFactoryJson              MemoryMemoryCategory = 21
	MemoryMemoryCategoryJsonUIControlTree                     MemoryMemoryCategory = 22
	MemoryMemoryCategoryJsonUIControlTreeControlElement       MemoryMemoryCategory = 23
	MemoryMemoryCategoryJsonUIControlTreePopulateDataBinding  MemoryMemoryCategory = 24
	MemoryMemoryCategoryJsonUIControlTreePopulateFocus        MemoryMemoryCategory = 25
	MemoryMemoryCategoryJsonUIControlTreePopulateLayout       MemoryMemoryCategory = 26
	MemoryMemoryCategoryJsonUIControlTreePopulateOther        MemoryMemoryCategory = 27
	MemoryMemoryCategoryJsonUIControlTreePopulateSprite       MemoryMemoryCategory = 28
	MemoryMemoryCategoryJsonUIControlTreePopulateText         MemoryMemoryCategory = 29
	MemoryMemoryCategoryJsonUIControlTreePopulateTTS          MemoryMemoryCategory = 30
	MemoryMemoryCategoryJsonUIControlTreeVisibility           MemoryMemoryCategory = 31
	MemoryMemoryCategoryJsonUICreateUI                        MemoryMemoryCategory = 32
	MemoryMemoryCategoryJsonUIDefs                            MemoryMemoryCategory = 33
	MemoryMemoryCategoryJsonUILayoutManager                   MemoryMemoryCategory = 34
	MemoryMemoryCategoryJsonUILayoutManagerRemoveDependencies MemoryMemoryCategory = 35
	MemoryMemoryCategoryJsonUILayoutManagerInitVariable       MemoryMemoryCategory = 36
	MemoryMemoryCategoryLanguages                             MemoryMemoryCategory = 37
	MemoryMemoryCategoryLevel                                 MemoryMemoryCategory = 38
	MemoryMemoryCategoryLevelStructures                       MemoryMemoryCategory = 39
	MemoryMemoryCategoryLevelChunk                            MemoryMemoryCategory = 40
	MemoryMemoryCategoryLevelChunkGen                         MemoryMemoryCategory = 41
	MemoryMemoryCategoryLevelChunkGenThreadLocal              MemoryMemoryCategory = 42
	MemoryMemoryCategoryLightVolumeManager                    MemoryMemoryCategory = 43
	MemoryMemoryCategoryNetwork                               MemoryMemoryCategory = 44
	MemoryMemoryCategoryMarketplace                           MemoryMemoryCategory = 45
	MemoryMemoryCategoryMaterialDragonCompiledDefinition      MemoryMemoryCategory = 46
	MemoryMemoryCategoryMaterialDragonMaterial                MemoryMemoryCategory = 47
	MemoryMemoryCategoryMaterialDragonResource                MemoryMemoryCategory = 48
	MemoryMemoryCategoryMaterialDragonUniformMap              MemoryMemoryCategory = 49
	MemoryMemoryCategoryMaterialRenderMaterial                MemoryMemoryCategory = 50
	MemoryMemoryCategoryMaterialRenderMaterialGroup           MemoryMemoryCategory = 51
	MemoryMemoryCategoryMaterialVariationManager              MemoryMemoryCategory = 52
	MemoryMemoryCategoryMolang                                MemoryMemoryCategory = 53
	MemoryMemoryCategoryOreUI                                 MemoryMemoryCategory = 54
	MemoryMemoryCategoryOreUIClient                           MemoryMemoryCategory = 55
	MemoryMemoryCategoryPersonaPieces                         MemoryMemoryCategory = 56
	MemoryMemoryCategoryPersonaAnimations                     MemoryMemoryCategory = 57
	MemoryMemoryCategoryPersonaTextures                       MemoryMemoryCategory = 58
	MemoryMemoryCategoryPersonaCharacters                     MemoryMemoryCategory = 59
	MemoryMemoryCategoryPersonaSkinPacks                      MemoryMemoryCategory = 60
	MemoryMemoryCategoryPersonaRepo                           MemoryMemoryCategory = 61
	MemoryMemoryCategoryPlayer                                MemoryMemoryCategory = 62
	MemoryMemoryCategoryRenderChunk                           MemoryMemoryCategory = 63
	MemoryMemoryCategoryRenderChunkIndexBuffer                MemoryMemoryCategory = 64
	MemoryMemoryCategoryRenderChunkVertexBuffer               MemoryMemoryCategory = 65
	MemoryMemoryCategoryRendering                             MemoryMemoryCategory = 66
	MemoryMemoryCategoryRenderingBgfxInit                     MemoryMemoryCategory = 67
	MemoryMemoryCategoryRenderingBgfxStartFrame               MemoryMemoryCategory = 68
	MemoryMemoryCategoryRenderingBlockTessellator             MemoryMemoryCategory = 69
	MemoryMemoryCategoryRenderingEndFrame                     MemoryMemoryCategory = 70
	MemoryMemoryCategoryRenderingGraphicsTasksInit            MemoryMemoryCategory = 71
	MemoryMemoryCategoryRenderingLibrary                      MemoryMemoryCategory = 72
	MemoryMemoryCategoryRenderingPolygonOperatorPool          MemoryMemoryCategory = 73
	MemoryMemoryCategoryRenderingPBRTextureData               MemoryMemoryCategory = 74
	MemoryMemoryCategoryRenderingRenderRegistry               MemoryMemoryCategory = 75
	MemoryMemoryCategoryRenderingSetup                        MemoryMemoryCategory = 76
	MemoryMemoryCategoryRenderingVertices                     MemoryMemoryCategory = 77
	MemoryMemoryCategoryRequestLog                            MemoryMemoryCategory = 78
	MemoryMemoryCategoryResourcePacks                         MemoryMemoryCategory = 79
	MemoryMemoryCategorySound                                 MemoryMemoryCategory = 80
	MemoryMemoryCategorySubChunkBiomeData                     MemoryMemoryCategory = 81
	MemoryMemoryCategorySubChunkBlockData                     MemoryMemoryCategory = 82
	MemoryMemoryCategorySubChunkLightData                     MemoryMemoryCategory = 83
	MemoryMemoryCategoryTextures                              MemoryMemoryCategory = 84
	MemoryMemoryCategoryWeatherRenderer                       MemoryMemoryCategory = 85
	MemoryMemoryCategoryWorldGenerator                        MemoryMemoryCategory = 86
	MemoryMemoryCategoryTasks                                 MemoryMemoryCategory = 87
	MemoryMemoryCategoryTest                                  MemoryMemoryCategory = 88
	MemoryMemoryCategoryTestLoadTestTags                      MemoryMemoryCategory = 89
	MemoryMemoryCategoryScripting                             MemoryMemoryCategory = 90
	MemoryMemoryCategoryScriptingRuntime                      MemoryMemoryCategory = 91
	MemoryMemoryCategoryScriptingContext                      MemoryMemoryCategory = 92
	MemoryMemoryCategoryScriptingContextBindingsMC            MemoryMemoryCategory = 93
	MemoryMemoryCategoryScriptingContextBindingsGT            MemoryMemoryCategory = 94
	MemoryMemoryCategoryScriptingContextRun                   MemoryMemoryCategory = 95
	MemoryMemoryCategoryDataDrivenUI                          MemoryMemoryCategory = 96
	MemoryMemoryCategoryDataDrivenUIDefs                      MemoryMemoryCategory = 97
	MemoryMemoryCategoryGameface                              MemoryMemoryCategory = 98
	MemoryMemoryCategoryGamefaceSystem                        MemoryMemoryCategory = 99
	MemoryMemoryCategoryGamefaceDOM                           MemoryMemoryCategory = 100
	MemoryMemoryCategoryGamefaceCSS                           MemoryMemoryCategory = 101
	MemoryMemoryCategoryGamefaceDisplay                       MemoryMemoryCategory = 102
	MemoryMemoryCategoryGamefaceTempAllocator                 MemoryMemoryCategory = 103
	MemoryMemoryCategoryGamefacePoolAllocator                 MemoryMemoryCategory = 104
	MemoryMemoryCategoryGamefaceDump                          MemoryMemoryCategory = 105
	MemoryMemoryCategoryGamefaceMedia                         MemoryMemoryCategory = 106
	MemoryMemoryCategoryGamefaceJSON                          MemoryMemoryCategory = 107
	MemoryMemoryCategoryGamefaceScriptEngine                  MemoryMemoryCategory = 108
	MemoryMemoryCategoryGamefaceScript                        MemoryMemoryCategory = 109
	MemoryMemoryCategoryGamefaceLayout                        MemoryMemoryCategory = 110
)

type MinecraftEventingAchievementIds uint8

const (
	MinecraftEventingAchievementIdsChestFullOfCobblestone          MinecraftEventingAchievementIds = 7
	MinecraftEventingAchievementIdsDiamondForYou                   MinecraftEventingAchievementIds = 10
	MinecraftEventingAchievementIdsIronBelly                       MinecraftEventingAchievementIds = 20
	MinecraftEventingAchievementIdsIronMan                         MinecraftEventingAchievementIds = 21
	MinecraftEventingAchievementIdsOnARail                         MinecraftEventingAchievementIds = 29
	MinecraftEventingAchievementIdsOverkill                        MinecraftEventingAchievementIds = 30
	MinecraftEventingAchievementIdsReturnToSender                  MinecraftEventingAchievementIds = 37
	MinecraftEventingAchievementIdsSniperDuel                      MinecraftEventingAchievementIds = 38
	MinecraftEventingAchievementIdsStayinFrosty                    MinecraftEventingAchievementIds = 39
	MinecraftEventingAchievementIdsTakeInventory                   MinecraftEventingAchievementIds = 40
	MinecraftEventingAchievementIdsMapRoom                         MinecraftEventingAchievementIds = 50
	MinecraftEventingAchievementIdsFreightStation                  MinecraftEventingAchievementIds = 52
	MinecraftEventingAchievementIdsSmeltEverything                 MinecraftEventingAchievementIds = 53
	MinecraftEventingAchievementIdsTasteOfYourOwnMedicine          MinecraftEventingAchievementIds = 54
	MinecraftEventingAchievementIdsWhenPigsFly                     MinecraftEventingAchievementIds = 56
	MinecraftEventingAchievementIdsInception                       MinecraftEventingAchievementIds = 58
	MinecraftEventingAchievementIdsArtificialSelection             MinecraftEventingAchievementIds = 60
	MinecraftEventingAchievementIdsFreeDiver                       MinecraftEventingAchievementIds = 61
	MinecraftEventingAchievementIdsSpawnTheWither                  MinecraftEventingAchievementIds = 62
	MinecraftEventingAchievementIdsBeaconator                      MinecraftEventingAchievementIds = 63
	MinecraftEventingAchievementIdsGreatView                       MinecraftEventingAchievementIds = 64
	MinecraftEventingAchievementIdsSuperSonic                      MinecraftEventingAchievementIds = 65
	MinecraftEventingAchievementIdsTheEndAgain                     MinecraftEventingAchievementIds = 66
	MinecraftEventingAchievementIdsTreasureHunter                  MinecraftEventingAchievementIds = 67
	MinecraftEventingAchievementIdsShootingStar                    MinecraftEventingAchievementIds = 68
	MinecraftEventingAchievementIdsFashionShow                     MinecraftEventingAchievementIds = 69
	MinecraftEventingAchievementIdsSelfPublishedAuthor             MinecraftEventingAchievementIds = 71
	MinecraftEventingAchievementIdsAlternativeFuel                 MinecraftEventingAchievementIds = 72
	MinecraftEventingAchievementIdsSleepWithTheFishes              MinecraftEventingAchievementIds = 73
	MinecraftEventingAchievementIdsCastaway                        MinecraftEventingAchievementIds = 74
	MinecraftEventingAchievementIdsImAMarineBiologist              MinecraftEventingAchievementIds = 75
	MinecraftEventingAchievementIdsSailThe7Seas                    MinecraftEventingAchievementIds = 76
	MinecraftEventingAchievementIdsMeGold                          MinecraftEventingAchievementIds = 77
	MinecraftEventingAchievementIdsAhoy                            MinecraftEventingAchievementIds = 78
	MinecraftEventingAchievementIdsAtlantis                        MinecraftEventingAchievementIds = 79
	MinecraftEventingAchievementIdsOnePickleTwoPickleSeaPickleFour MinecraftEventingAchievementIds = 80
	MinecraftEventingAchievementIdsDoaBarrelRoll                   MinecraftEventingAchievementIds = 81
	MinecraftEventingAchievementIdsMoskstraumen                    MinecraftEventingAchievementIds = 82
	MinecraftEventingAchievementIdsEcholocation                    MinecraftEventingAchievementIds = 83
	MinecraftEventingAchievementIdsWhereHaveYouBeen                MinecraftEventingAchievementIds = 84
	MinecraftEventingAchievementIdsTopOfTheWorld                   MinecraftEventingAchievementIds = 85
	MinecraftEventingAchievementIdsFruitOnTheLoom                  MinecraftEventingAchievementIds = 86
	MinecraftEventingAchievementIdsSoundTheAlarm                   MinecraftEventingAchievementIds = 87
	MinecraftEventingAchievementIdsBuyLowSellHigh                  MinecraftEventingAchievementIds = 88
	MinecraftEventingAchievementIdsDisenchanted                    MinecraftEventingAchievementIds = 89
	MinecraftEventingAchievementIdsTimeForStew                     MinecraftEventingAchievementIds = 90
	MinecraftEventingAchievementIdsBeeOurGuest                     MinecraftEventingAchievementIds = 91
	MinecraftEventingAchievementIdsTotalBeeLocation                MinecraftEventingAchievementIds = 92
	MinecraftEventingAchievementIdsStickySituation                 MinecraftEventingAchievementIds = 93
	MinecraftEventingAchievementIdsCoverMeInDebris                 MinecraftEventingAchievementIds = 94
	MinecraftEventingAchievementIdsFloatYourGoat                   MinecraftEventingAchievementIds = 95
	MinecraftEventingAchievementIdsFriend                          MinecraftEventingAchievementIds = 96
	MinecraftEventingAchievementIdsWaxOnWaxOff                     MinecraftEventingAchievementIds = 97
	MinecraftEventingAchievementIdsStriderRiddenInLavaInOverworld  MinecraftEventingAchievementIds = 98
	MinecraftEventingAchievementIdsGoatHornAcquired                MinecraftEventingAchievementIds = 99
	MinecraftEventingAchievementIdsJukeboxUsedInMeadows            MinecraftEventingAchievementIds = 100
	MinecraftEventingAchievementIdsTradedAtWorldHeight             MinecraftEventingAchievementIds = 101
	MinecraftEventingAchievementIdsSurvivedFallFromWorldHeight     MinecraftEventingAchievementIds = 102
	MinecraftEventingAchievementIdsSneakCloseToSculkSensor         MinecraftEventingAchievementIds = 103
	MinecraftEventingAchievementIdsItSpreads                       MinecraftEventingAchievementIds = 104
	MinecraftEventingAchievementIdsBirthdaySong                    MinecraftEventingAchievementIds = 105
	MinecraftEventingAchievementIdsWithOurPowersCombined           MinecraftEventingAchievementIds = 106
	MinecraftEventingAchievementIdsPlantingThePast                 MinecraftEventingAchievementIds = 107
	MinecraftEventingAchievementIdsCarefulRestoration              MinecraftEventingAchievementIds = 108
	MinecraftEventingAchievementIdsRevaulting                      MinecraftEventingAchievementIds = 109
	MinecraftEventingAchievementIdsCraftersCraftingCrafters        MinecraftEventingAchievementIds = 110
	MinecraftEventingAchievementIdsWhoNeedsRockets                 MinecraftEventingAchievementIds = 111
	MinecraftEventingAchievementIdsOverOverkill                    MinecraftEventingAchievementIds = 112
	MinecraftEventingAchievementIdsHeartTransplanter               MinecraftEventingAchievementIds = 113
	MinecraftEventingAchievementIdsStayHydrated                    MinecraftEventingAchievementIds = 114
	MinecraftEventingAchievementIdsMobKabob                        MinecraftEventingAchievementIds = 115
	MinecraftEventingAchievementIdsAdventuringTime                 MinecraftEventingAchievementIds = 116
	MinecraftEventingAchievementIdsUhOh                            MinecraftEventingAchievementIds = 117
	MinecraftEventingAchievementIdsGettingWood                     MinecraftEventingAchievementIds = 118
	MinecraftEventingAchievementIdsBenchMaking                     MinecraftEventingAchievementIds = 119
	MinecraftEventingAchievementIdsTimeToMine                      MinecraftEventingAchievementIds = 120
	MinecraftEventingAchievementIdsHotTopic                        MinecraftEventingAchievementIds = 121
	MinecraftEventingAchievementIdsAcquireHardware                 MinecraftEventingAchievementIds = 122
	MinecraftEventingAchievementIdsGettingAnUpgrade                MinecraftEventingAchievementIds = 123
	MinecraftEventingAchievementIdsMonsterHunter                   MinecraftEventingAchievementIds = 124
	MinecraftEventingAchievementIdsDiamonds                        MinecraftEventingAchievementIds = 125
	MinecraftEventingAchievementIdsPlethoraOfCats                  MinecraftEventingAchievementIds = 126
)

type MinecraftEventingInteractionType uint8

const (
	MinecraftEventingInteractionTypeBreeding   MinecraftEventingInteractionType = 1
	MinecraftEventingInteractionTypeTaming     MinecraftEventingInteractionType = 2
	MinecraftEventingInteractionTypeCuring     MinecraftEventingInteractionType = 3
	MinecraftEventingInteractionTypeCrafted    MinecraftEventingInteractionType = 4
	MinecraftEventingInteractionTypeShearing   MinecraftEventingInteractionType = 5
	MinecraftEventingInteractionTypeMilking    MinecraftEventingInteractionType = 6
	MinecraftEventingInteractionTypeTrading    MinecraftEventingInteractionType = 7
	MinecraftEventingInteractionTypeFeeding    MinecraftEventingInteractionType = 8
	MinecraftEventingInteractionTypeIgniting   MinecraftEventingInteractionType = 9
	MinecraftEventingInteractionTypeColoring   MinecraftEventingInteractionType = 10
	MinecraftEventingInteractionTypeNaming     MinecraftEventingInteractionType = 11
	MinecraftEventingInteractionTypeLeashing   MinecraftEventingInteractionType = 12
	MinecraftEventingInteractionTypeUnleashing MinecraftEventingInteractionType = 13
	MinecraftEventingInteractionTypePetSleep   MinecraftEventingInteractionType = 14
	MinecraftEventingInteractionTypeTrusting   MinecraftEventingInteractionType = 15
	MinecraftEventingInteractionTypeCommanding MinecraftEventingInteractionType = 16
	MinecraftEventingInteractionTypeEquipping  MinecraftEventingInteractionType = 17
)

type MinecraftEventingPOIBlockInteractionType uint8

const (
	MinecraftEventingPOIBlockInteractionTypeNone                MinecraftEventingPOIBlockInteractionType = 0
	MinecraftEventingPOIBlockInteractionTypeExtend              MinecraftEventingPOIBlockInteractionType = 1
	MinecraftEventingPOIBlockInteractionTypeClone               MinecraftEventingPOIBlockInteractionType = 2
	MinecraftEventingPOIBlockInteractionTypeLock                MinecraftEventingPOIBlockInteractionType = 3
	MinecraftEventingPOIBlockInteractionTypeCreate              MinecraftEventingPOIBlockInteractionType = 4
	MinecraftEventingPOIBlockInteractionTypeCreateLocator       MinecraftEventingPOIBlockInteractionType = 5
	MinecraftEventingPOIBlockInteractionTypeRename              MinecraftEventingPOIBlockInteractionType = 6
	MinecraftEventingPOIBlockInteractionTypeItemPlaced          MinecraftEventingPOIBlockInteractionType = 7
	MinecraftEventingPOIBlockInteractionTypeItemRemoved         MinecraftEventingPOIBlockInteractionType = 8
	MinecraftEventingPOIBlockInteractionTypeCooking             MinecraftEventingPOIBlockInteractionType = 9
	MinecraftEventingPOIBlockInteractionTypeDousing             MinecraftEventingPOIBlockInteractionType = 10
	MinecraftEventingPOIBlockInteractionTypeLighting            MinecraftEventingPOIBlockInteractionType = 11
	MinecraftEventingPOIBlockInteractionTypeHaystack            MinecraftEventingPOIBlockInteractionType = 12
	MinecraftEventingPOIBlockInteractionTypeFilled              MinecraftEventingPOIBlockInteractionType = 13
	MinecraftEventingPOIBlockInteractionTypeEmptied             MinecraftEventingPOIBlockInteractionType = 14
	MinecraftEventingPOIBlockInteractionTypeAddDye              MinecraftEventingPOIBlockInteractionType = 15
	MinecraftEventingPOIBlockInteractionTypeDyeItem             MinecraftEventingPOIBlockInteractionType = 16
	MinecraftEventingPOIBlockInteractionTypeClearItem           MinecraftEventingPOIBlockInteractionType = 17
	MinecraftEventingPOIBlockInteractionTypeEnchantArrow        MinecraftEventingPOIBlockInteractionType = 18
	MinecraftEventingPOIBlockInteractionTypeCompostItemPlaced   MinecraftEventingPOIBlockInteractionType = 19
	MinecraftEventingPOIBlockInteractionTypeRecoveredBonemeal   MinecraftEventingPOIBlockInteractionType = 20
	MinecraftEventingPOIBlockInteractionTypeBookPlaced          MinecraftEventingPOIBlockInteractionType = 21
	MinecraftEventingPOIBlockInteractionTypeBookOpened          MinecraftEventingPOIBlockInteractionType = 22
	MinecraftEventingPOIBlockInteractionTypeDisenchant          MinecraftEventingPOIBlockInteractionType = 23
	MinecraftEventingPOIBlockInteractionTypeRepair              MinecraftEventingPOIBlockInteractionType = 24
	MinecraftEventingPOIBlockInteractionTypeDisenchantAndRepair MinecraftEventingPOIBlockInteractionType = 25
)

type Mirror uint8

const (
	MirrorNone Mirror = 0
	MirrorX    Mirror = 1
	MirrorZ    Mirror = 2
	MirrorXZ   Mirror = 3
)

type MoLangVersion int16

const (
	MoLangVersionInvalid                                MoLangVersion = -1
	MoLangVersionBeforeVersioning                       MoLangVersion = 0
	MoLangVersionInitial                                MoLangVersion = 1
	MoLangVersionFixedItemRemainingUseDurationQuery     MoLangVersion = 2
	MoLangVersionExpressionErrorMessages                MoLangVersion = 3
	MoLangVersionUnexpectedOperatorErrors               MoLangVersion = 4
	MoLangVersionConditionalOperatorAssociativity       MoLangVersion = 5
	MoLangVersionComparisonAndLogicalOperatorPrecedence MoLangVersion = 6
	MoLangVersionDivideByNegativeValue                  MoLangVersion = 7
	MoLangVersionFixedCapeFlapAmountQuery               MoLangVersion = 8
	MoLangVersionQueryBlockPropertyRenamedToState       MoLangVersion = 9
	MoLangVersionDeprecateOldBlockQueryNames            MoLangVersion = 10
	MoLangVersionDeprecatedSnifferAndCamelQueries       MoLangVersion = 11
	MoLangVersionLeafSupportingInFirstSolidBlockBelow   MoLangVersion = 12
	MoLangVersionLatest                                 MoLangVersion = 13
	MoLangVersionNumValidVersions                       MoLangVersion = 14
)

type MobEffectEvent uint8

const (
	MobEffectEventInvalid MobEffectEvent = 0
	MobEffectEventAdd     MobEffectEvent = 1
	MobEffectEventUpdate  MobEffectEvent = 2
	MobEffectEventRemove  MobEffectEvent = 3
)

type ModalFormCancelReason uint8

const (
	ModalFormCancelReasonUserClosed ModalFormCancelReason = 0
	ModalFormCancelReasonUserBusy   ModalFormCancelReason = 1
)

type MovementEffectType int32

const (
	MovementEffectTypeGLIDEBOOST   MovementEffectType = 0
	MovementEffectTypeDOLPHINBOOST MovementEffectType = 1
	MovementEffectTypeGEYSERBOOST  MovementEffectType = 2
)

type MultiplayerSettingsPacketType int32

const (
	MultiplayerSettingsPacketTypeEnableMultiplayer  MultiplayerSettingsPacketType = 0
	MultiplayerSettingsPacketTypeDisableMultiplayer MultiplayerSettingsPacketType = 1
	MultiplayerSettingsPacketTypeRefreshJoincode    MultiplayerSettingsPacketType = 2
)

type NewInteractionModel int32

const (
	NewInteractionModelTouch     NewInteractionModel = 0
	NewInteractionModelCrosshair NewInteractionModel = 1
	NewInteractionModelClassic   NewInteractionModel = 2
	NewInteractionModelCount     NewInteractionModel = 3
)

type NpcDialogueNpcDialogueActionType int32

const (
	NpcDialogueNpcDialogueActionTypeOpen  NpcDialogueNpcDialogueActionType = 0
	NpcDialogueNpcDialogueActionTypeClose NpcDialogueNpcDialogueActionType = 1
)

type NpcRequestRequestType uint8

const (
	NpcRequestRequestTypeSetActions             NpcRequestRequestType = 0
	NpcRequestRequestTypeExecuteAction          NpcRequestRequestType = 1
	NpcRequestRequestTypeExecuteClosingCommands NpcRequestRequestType = 2
	NpcRequestRequestTypeSetName                NpcRequestRequestType = 3
	NpcRequestRequestTypeSetSkin                NpcRequestRequestType = 4
	NpcRequestRequestTypeSetInteractText        NpcRequestRequestType = 5
	NpcRequestRequestTypeExecuteOpeningCommands NpcRequestRequestType = 6
)

type PacketCompressionAlgorithm uint16

const (
	PacketCompressionAlgorithmZLib   PacketCompressionAlgorithm = 0
	PacketCompressionAlgorithmSnappy PacketCompressionAlgorithm = 1
	PacketCompressionAlgorithmNone   PacketCompressionAlgorithm = 65535
)

type PacketViolationSeverity int32

const (
	PacketViolationSeverityUnknown               PacketViolationSeverity = -1
	PacketViolationSeverityWarning               PacketViolationSeverity = 0
	PacketViolationSeverityFinalWarning          PacketViolationSeverity = 1
	PacketViolationSeverityTerminatingConnection PacketViolationSeverity = 2
)

type PacketViolationType int32

const (
	PacketViolationTypeUnknown         PacketViolationType = -1
	PacketViolationTypePacketMalformed PacketViolationType = 0
)

type PersonaAnimatedTextureType uint32

const (
	PersonaAnimatedTextureTypeFace        PersonaAnimatedTextureType = 1
	PersonaAnimatedTextureTypeBody32x32   PersonaAnimatedTextureType = 2
	PersonaAnimatedTextureTypeBody128x128 PersonaAnimatedTextureType = 3
)

type PersonaAnimationExpression uint32

const (
	PersonaAnimationExpressionLinear   PersonaAnimationExpression = 0
	PersonaAnimationExpressionBlinking PersonaAnimationExpression = 1
)

type PersonaArmSizeType uint8

const (
	PersonaArmSizeTypeSlim PersonaArmSizeType = 0
	PersonaArmSizeTypeWide PersonaArmSizeType = 1
)

type PersonaPieceType uint32

const (
	PersonaPieceTypeSkeleton      PersonaPieceType = 1
	PersonaPieceTypeBody          PersonaPieceType = 2
	PersonaPieceTypeSkin          PersonaPieceType = 3
	PersonaPieceTypeBottom        PersonaPieceType = 4
	PersonaPieceTypeFeet          PersonaPieceType = 5
	PersonaPieceTypeDress         PersonaPieceType = 6
	PersonaPieceTypeTop           PersonaPieceType = 7
	PersonaPieceTypeHighPants     PersonaPieceType = 8
	PersonaPieceTypeHands         PersonaPieceType = 9
	PersonaPieceTypeOuterwear     PersonaPieceType = 10
	PersonaPieceTypeFacialHair    PersonaPieceType = 11
	PersonaPieceTypeMouth         PersonaPieceType = 12
	PersonaPieceTypeEyes          PersonaPieceType = 13
	PersonaPieceTypeHair          PersonaPieceType = 14
	PersonaPieceTypeHood          PersonaPieceType = 15
	PersonaPieceTypeBack          PersonaPieceType = 16
	PersonaPieceTypeFaceAccessory PersonaPieceType = 17
	PersonaPieceTypeHead          PersonaPieceType = 18
	PersonaPieceTypeLegs          PersonaPieceType = 19
	PersonaPieceTypeLeftLeg       PersonaPieceType = 20
	PersonaPieceTypeRightLeg      PersonaPieceType = 21
	PersonaPieceTypeArms          PersonaPieceType = 22
	PersonaPieceTypeLeftArm       PersonaPieceType = 23
	PersonaPieceTypeRightArm      PersonaPieceType = 24
	PersonaPieceTypeCapes         PersonaPieceType = 25
	PersonaPieceTypeClassicSkin   PersonaPieceType = 26
	PersonaPieceTypeEmote         PersonaPieceType = 27
)

type PhotoType uint8

const (
	PhotoTypePortfolio PhotoType = 0
	PhotoTypePhotoItem PhotoType = 1
	PhotoTypeBook      PhotoType = 2
)

type PlayStatusType int32

const (
	PlayStatusTypeLoginSuccess                             PlayStatusType = 0
	PlayStatusTypeLoginFailedClientOld                     PlayStatusType = 1
	PlayStatusTypeLoginFailedServerOld                     PlayStatusType = 2
	PlayStatusTypePlayerSpawn                              PlayStatusType = 3
	PlayStatusTypeLoginFailedInvalidTenant                 PlayStatusType = 4
	PlayStatusTypeLoginFailedEditionMismatchEduToVanilla   PlayStatusType = 5
	PlayStatusTypeLoginFailedEditionMismatchVanillaToEdu   PlayStatusType = 6
	PlayStatusTypeLoginFailedServerFullSubClient           PlayStatusType = 7
	PlayStatusTypeLoginFailedEditorMismatchEditorToVanilla PlayStatusType = 8
	PlayStatusTypeLoginFailedEditorMismatchVanillaToEditor PlayStatusType = 9
)

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

type PlayerAuthInputInputData int32

const (
	PlayerAuthInputInputDataAscend                          PlayerAuthInputInputData = 0
	PlayerAuthInputInputDataDescend                         PlayerAuthInputInputData = 1
	PlayerAuthInputInputDataNorthJump                       PlayerAuthInputInputData = 2
	PlayerAuthInputInputDataJumpDown                        PlayerAuthInputInputData = 3
	PlayerAuthInputInputDataSprintDown                      PlayerAuthInputInputData = 4
	PlayerAuthInputInputDataChangeHeight                    PlayerAuthInputInputData = 5
	PlayerAuthInputInputDataJumping                         PlayerAuthInputInputData = 6
	PlayerAuthInputInputDataAutoJumpingInWater              PlayerAuthInputInputData = 7
	PlayerAuthInputInputDataSneaking                        PlayerAuthInputInputData = 8
	PlayerAuthInputInputDataSneakDown                       PlayerAuthInputInputData = 9
	PlayerAuthInputInputDataUp                              PlayerAuthInputInputData = 10
	PlayerAuthInputInputDataDown                            PlayerAuthInputInputData = 11
	PlayerAuthInputInputDataLeft                            PlayerAuthInputInputData = 12
	PlayerAuthInputInputDataRight                           PlayerAuthInputInputData = 13
	PlayerAuthInputInputDataUpLeft                          PlayerAuthInputInputData = 14
	PlayerAuthInputInputDataUpRight                         PlayerAuthInputInputData = 15
	PlayerAuthInputInputDataWantUp                          PlayerAuthInputInputData = 16
	PlayerAuthInputInputDataWantDown                        PlayerAuthInputInputData = 17
	PlayerAuthInputInputDataWantDownSlow                    PlayerAuthInputInputData = 18
	PlayerAuthInputInputDataWantUpSlow                      PlayerAuthInputInputData = 19
	PlayerAuthInputInputDataSprinting                       PlayerAuthInputInputData = 20
	PlayerAuthInputInputDataAscendBlock                     PlayerAuthInputInputData = 21
	PlayerAuthInputInputDataDescendBlock                    PlayerAuthInputInputData = 22
	PlayerAuthInputInputDataSneakToggleDown                 PlayerAuthInputInputData = 23
	PlayerAuthInputInputDataPersistSneak                    PlayerAuthInputInputData = 24
	PlayerAuthInputInputDataStartSprinting                  PlayerAuthInputInputData = 25
	PlayerAuthInputInputDataStopSprinting                   PlayerAuthInputInputData = 26
	PlayerAuthInputInputDataStartSneaking                   PlayerAuthInputInputData = 27
	PlayerAuthInputInputDataStopSneaking                    PlayerAuthInputInputData = 28
	PlayerAuthInputInputDataStartSwimming                   PlayerAuthInputInputData = 29
	PlayerAuthInputInputDataStopSwimming                    PlayerAuthInputInputData = 30
	PlayerAuthInputInputDataStartJumping                    PlayerAuthInputInputData = 31
	PlayerAuthInputInputDataStartGliding                    PlayerAuthInputInputData = 32
	PlayerAuthInputInputDataStopGliding                     PlayerAuthInputInputData = 33
	PlayerAuthInputInputDataPerformItemInteraction          PlayerAuthInputInputData = 34
	PlayerAuthInputInputDataPerformBlockActions             PlayerAuthInputInputData = 35
	PlayerAuthInputInputDataPerformItemStackRequest         PlayerAuthInputInputData = 36
	PlayerAuthInputInputDataHandledTeleport                 PlayerAuthInputInputData = 37
	PlayerAuthInputInputDataEmoting                         PlayerAuthInputInputData = 38
	PlayerAuthInputInputDataMissedSwing                     PlayerAuthInputInputData = 39
	PlayerAuthInputInputDataStartCrawling                   PlayerAuthInputInputData = 40
	PlayerAuthInputInputDataStopCrawling                    PlayerAuthInputInputData = 41
	PlayerAuthInputInputDataStartFlying                     PlayerAuthInputInputData = 42
	PlayerAuthInputInputDataStopFlying                      PlayerAuthInputInputData = 43
	PlayerAuthInputInputDataClientAckServerData             PlayerAuthInputInputData = 44
	PlayerAuthInputInputDataIsInClientPredictedVehicle      PlayerAuthInputInputData = 45
	PlayerAuthInputInputDataPaddlingLeft                    PlayerAuthInputInputData = 46
	PlayerAuthInputInputDataPaddlingRight                   PlayerAuthInputInputData = 47
	PlayerAuthInputInputDataBlockBreakingDelayEnabled       PlayerAuthInputInputData = 48
	PlayerAuthInputInputDataHorizontalCollision             PlayerAuthInputInputData = 49
	PlayerAuthInputInputDataVerticalCollision               PlayerAuthInputInputData = 50
	PlayerAuthInputInputDataDownLeft                        PlayerAuthInputInputData = 51
	PlayerAuthInputInputDataDownRight                       PlayerAuthInputInputData = 52
	PlayerAuthInputInputDataStartUsingItem                  PlayerAuthInputInputData = 53
	PlayerAuthInputInputDataIsCameraRelativeMovementEnabled PlayerAuthInputInputData = 54
	PlayerAuthInputInputDataIsRotControlledByMoveDirection  PlayerAuthInputInputData = 55
	PlayerAuthInputInputDataStartSpinAttack                 PlayerAuthInputInputData = 56
	PlayerAuthInputInputDataStopSpinAttack                  PlayerAuthInputInputData = 57
	PlayerAuthInputInputDataIsHotbarOnlyTouch               PlayerAuthInputInputData = 58
	PlayerAuthInputInputDataJumpReleasedRaw                 PlayerAuthInputInputData = 59
	PlayerAuthInputInputDataJumpPressedRaw                  PlayerAuthInputInputData = 60
	PlayerAuthInputInputDataJumpCurrentRaw                  PlayerAuthInputInputData = 61
	PlayerAuthInputInputDataSneakReleasedRaw                PlayerAuthInputInputData = 62
	PlayerAuthInputInputDataSneakPressedRaw                 PlayerAuthInputInputData = 63
	PlayerAuthInputInputDataSneakCurrentRaw                 PlayerAuthInputInputData = 64
	PlayerAuthInputInputDataInternalUpdate                  PlayerAuthInputInputData = 65
)

type PlayerLocationType int32

const (
	PlayerLocationTypePLAYERLOCATIONCOORDINATES PlayerLocationType = 0
)

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

type PositionTrackingDBClientRequestAction uint8

const (
	PositionTrackingDBClientRequestActionQuery PositionTrackingDBClientRequestAction = 0
)

type PositionTrackingDBServerBroadcastAction uint8

const (
	PositionTrackingDBServerBroadcastActionUpdate   PositionTrackingDBServerBroadcastAction = 0
	PositionTrackingDBServerBroadcastActionDestroy  PositionTrackingDBServerBroadcastAction = 1
	PositionTrackingDBServerBroadcastActionNotFound PositionTrackingDBServerBroadcastAction = 2
)

type RandomDistributionType int32

const (
	RandomDistributionTypeSingleValued    RandomDistributionType = 0
	RandomDistributionTypeUniform         RandomDistributionType = 1
	RandomDistributionTypeGaussian        RandomDistributionType = 2
	RandomDistributionTypeInverseGaussian RandomDistributionType = 3
	RandomDistributionTypeFixedGrid       RandomDistributionType = 4
	RandomDistributionTypeJitteredGrid    RandomDistributionType = 5
	RandomDistributionTypeTriangle        RandomDistributionType = 6
)

type RecipeUnlockingRequirementUnlockingContext int32

const (
	RecipeUnlockingRequirementUnlockingContextNone               RecipeUnlockingRequirementUnlockingContext = 0
	RecipeUnlockingRequirementUnlockingContextAlwaysUnlocked     RecipeUnlockingRequirementUnlockingContext = 1
	RecipeUnlockingRequirementUnlockingContextPlayerInWater      RecipeUnlockingRequirementUnlockingContext = 2
	RecipeUnlockingRequirementUnlockingContextPlayerHasManyItems RecipeUnlockingRequirementUnlockingContext = 3
)

type RequestAbilityType uint8

const (
	RequestAbilityTypeUnset RequestAbilityType = 0
	RequestAbilityTypeBool  RequestAbilityType = 1
	RequestAbilityTypeFloat RequestAbilityType = 2
)

type RewindType uint8

const (
	RewindTypePlayer  RewindType = 0
	RewindTypeVehicle RewindType = 1
)

type Rotation uint8

const (
	RotationNone      Rotation = 0
	RotationRotate90  Rotation = 1
	RotationRotate180 Rotation = 2
	RotationRotate270 Rotation = 3
)

type ScoreboardIdentityPacketType uint8

const (
	ScoreboardIdentityPacketTypeUpdate ScoreboardIdentityPacketType = 0
	ScoreboardIdentityPacketTypeRemove ScoreboardIdentityPacketType = 1
)

type ScriptModuleMinecraftScriptPrimitiveShapeType uint8

const (
	ScriptModuleMinecraftScriptPrimitiveShapeTypeLine      ScriptModuleMinecraftScriptPrimitiveShapeType = 0
	ScriptModuleMinecraftScriptPrimitiveShapeTypeBox       ScriptModuleMinecraftScriptPrimitiveShapeType = 1
	ScriptModuleMinecraftScriptPrimitiveShapeTypeSphere    ScriptModuleMinecraftScriptPrimitiveShapeType = 2
	ScriptModuleMinecraftScriptPrimitiveShapeTypeCircle    ScriptModuleMinecraftScriptPrimitiveShapeType = 3
	ScriptModuleMinecraftScriptPrimitiveShapeTypeText      ScriptModuleMinecraftScriptPrimitiveShapeType = 4
	ScriptModuleMinecraftScriptPrimitiveShapeTypeArrow     ScriptModuleMinecraftScriptPrimitiveShapeType = 5
	ScriptModuleMinecraftScriptPrimitiveShapeTypeCylinder  ScriptModuleMinecraftScriptPrimitiveShapeType = 6
	ScriptModuleMinecraftScriptPrimitiveShapeTypePyramid   ScriptModuleMinecraftScriptPrimitiveShapeType = 7
	ScriptModuleMinecraftScriptPrimitiveShapeTypeEllipsoid ScriptModuleMinecraftScriptPrimitiveShapeType = 8
	ScriptModuleMinecraftScriptPrimitiveShapeTypeCone      ScriptModuleMinecraftScriptPrimitiveShapeType = 9
)

type ServerEditorConnectionPolicy int32

const (
	ServerEditorConnectionPolicyMatchWorldType ServerEditorConnectionPolicy = 0
	ServerEditorConnectionPolicyEditorOnly     ServerEditorConnectionPolicy = 1
	ServerEditorConnectionPolicyVanillaOnly    ServerEditorConnectionPolicy = 2
	ServerEditorConnectionPolicyMixed          ServerEditorConnectionPolicy = 3
)

type ServerWaypointGroupAction uint8

const (
	ServerWaypointGroupActionNone   ServerWaypointGroupAction = 0
	ServerWaypointGroupActionAdd    ServerWaypointGroupAction = 1
	ServerWaypointGroupActionRemove ServerWaypointGroupAction = 2
	ServerWaypointGroupActionUpdate ServerWaypointGroupAction = 3
)

type ServerboundLoadingScreenPacketType int32

const (
	ServerboundLoadingScreenPacketTypeStartLoadingScreen ServerboundLoadingScreenPacketType = 1
	ServerboundLoadingScreenPacketTypeEndLoadingScreen   ServerboundLoadingScreenPacketType = 2
)

type SetTitleTitleType int32

const (
	SetTitleTitleTypeClear               SetTitleTitleType = 0
	SetTitleTitleTypeReset               SetTitleTitleType = 1
	SetTitleTitleTypeTitle               SetTitleTitleType = 2
	SetTitleTitleTypeSubtitle            SetTitleTitleType = 3
	SetTitleTitleTypeActionbar           SetTitleTitleType = 4
	SetTitleTitleTypeTimes               SetTitleTitleType = 5
	SetTitleTitleTypeTitleTextObject     SetTitleTitleType = 6
	SetTitleTitleTypeSubtitleTextObject  SetTitleTitleType = 7
	SetTitleTitleTypeActionbarTextObject SetTitleTitleType = 8
)

type ShowStoreOfferRedirectType uint8

const (
	ShowStoreOfferRedirectTypeMarketplaceOffer     ShowStoreOfferRedirectType = 0
	ShowStoreOfferRedirectTypeDressingRoomOffer    ShowStoreOfferRedirectType = 1
	ShowStoreOfferRedirectTypeThirdPartyServerPage ShowStoreOfferRedirectType = 2
)

type SimpleEventSubtype uint16

const (
	SimpleEventSubtypeUninitializedSubtype        SimpleEventSubtype = 0
	SimpleEventSubtypeEnableCommands              SimpleEventSubtype = 1
	SimpleEventSubtypeDisableCommands             SimpleEventSubtype = 2
	SimpleEventSubtypeUnlockWorldTemplateSettings SimpleEventSubtype = 3
)

type SimulationTypeType uint8

const (
	SimulationTypeTypeGame    SimulationTypeType = 0
	SimulationTypeTypeEditor  SimulationTypeType = 1
	SimulationTypeTypeTest    SimulationTypeType = 2
	SimulationTypeTypeINVALID SimulationTypeType = 3
)

type SocialGamePublishSetting int32

const (
	SocialGamePublishSettingNoMultiPlay      SocialGamePublishSetting = 0
	SocialGamePublishSettingInviteOnly       SocialGamePublishSetting = 1
	SocialGamePublishSettingFriendsOnly      SocialGamePublishSetting = 2
	SocialGamePublishSettingFriendsOfFriends SocialGamePublishSetting = 3
	SocialGamePublishSettingPublic           SocialGamePublishSetting = 4
)

type SoftEnumUpdateType uint8

const (
	SoftEnumUpdateTypeAdd     SoftEnumUpdateType = 0
	SoftEnumUpdateTypeRemove  SoftEnumUpdateType = 1
	SoftEnumUpdateTypeReplace SoftEnumUpdateType = 2
)

type SpawnBiomeType int16

const (
	SpawnBiomeTypeDefault     SpawnBiomeType = 0
	SpawnBiomeTypeUserDefined SpawnBiomeType = 1
)

type SpawnPositionType int32

const (
	SpawnPositionTypePlayerRespawn SpawnPositionType = 0
	SpawnPositionTypeWorldSpawn    SpawnPositionType = 1
)

type StructureBlockType int32

const (
	StructureBlockTypeData    StructureBlockType = 0
	StructureBlockTypeSave    StructureBlockType = 1
	StructureBlockTypeLoad    StructureBlockType = 2
	StructureBlockTypeCorner  StructureBlockType = 3
	StructureBlockTypeInvalid StructureBlockType = 4
	StructureBlockTypeExport  StructureBlockType = 5
)

type StructureRedstoneSaveMode uint8

const (
	StructureRedstoneSaveModeSavesToMemory StructureRedstoneSaveMode = 0
	StructureRedstoneSaveModeSavesToDisk   StructureRedstoneSaveMode = 1
)

type StructureTemplateRequestOperation uint8

const (
	StructureTemplateRequestOperationNone                StructureTemplateRequestOperation = 0
	StructureTemplateRequestOperationExportFromSaveMode  StructureTemplateRequestOperation = 1
	StructureTemplateRequestOperationExportFromLoadMode  StructureTemplateRequestOperation = 2
	StructureTemplateRequestOperationQuerySavedStructure StructureTemplateRequestOperation = 3
)

type StructureTemplateResponseType uint8

const (
	StructureTemplateResponseTypeNone   StructureTemplateResponseType = 0
	StructureTemplateResponseTypeExport StructureTemplateResponseType = 1
	StructureTemplateResponseTypeQuery  StructureTemplateResponseType = 2
)

type SubChunkHeightMapDataType uint8

const (
	SubChunkHeightMapDataTypeNoData     SubChunkHeightMapDataType = 0
	SubChunkHeightMapDataTypeHasData    SubChunkHeightMapDataType = 1
	SubChunkHeightMapDataTypeAllTooHigh SubChunkHeightMapDataType = 2
	SubChunkHeightMapDataTypeAllTooLow  SubChunkHeightMapDataType = 3
)

type SubChunkSubChunkRequestResult uint8

const (
	SubChunkSubChunkRequestResultSuccess               SubChunkSubChunkRequestResult = 1
	SubChunkSubChunkRequestResultLevelChunkDoesntExist SubChunkSubChunkRequestResult = 2
	SubChunkSubChunkRequestResultWrongDimension        SubChunkSubChunkRequestResult = 3
	SubChunkSubChunkRequestResultPlayerDoesntExist     SubChunkSubChunkRequestResult = 4
	SubChunkSubChunkRequestResultIndexOutOfBounds      SubChunkSubChunkRequestResult = 5
	SubChunkSubChunkRequestResultSuccessAllAir         SubChunkSubChunkRequestResult = 6
)

type TextProcessingEventOrigin int32

const (
	TextProcessingEventOriginUnknown            TextProcessingEventOrigin = -1
	TextProcessingEventOriginServerChatPublic   TextProcessingEventOrigin = 0
	TextProcessingEventOriginServerChatWhisper  TextProcessingEventOrigin = 1
	TextProcessingEventOriginSignText           TextProcessingEventOrigin = 2
	TextProcessingEventOriginAnvilText          TextProcessingEventOrigin = 3
	TextProcessingEventOriginBookAndQuillText   TextProcessingEventOrigin = 4
	TextProcessingEventOriginCommandBlockText   TextProcessingEventOrigin = 5
	TextProcessingEventOriginBlockActorDataText TextProcessingEventOrigin = 6
	TextProcessingEventOriginJoinEventText      TextProcessingEventOrigin = 7
	TextProcessingEventOriginLeaveEventText     TextProcessingEventOrigin = 8
	TextProcessingEventOriginSlashCommandChat   TextProcessingEventOrigin = 9
	TextProcessingEventOriginCartographyText    TextProcessingEventOrigin = 10
	TextProcessingEventOriginKickCommand        TextProcessingEventOrigin = 11
	TextProcessingEventOriginTitleCommand       TextProcessingEventOrigin = 12
	TextProcessingEventOriginSummonCommand      TextProcessingEventOrigin = 13
	TextProcessingEventOriginServerForm         TextProcessingEventOrigin = 14
	TextProcessingEventOriginDataDrivenUI       TextProcessingEventOrigin = 15
)

type UnlockedRecipesPacketType uint32

const (
	UnlockedRecipesPacketTypeEmpty                    UnlockedRecipesPacketType = 0
	UnlockedRecipesPacketTypeInitiallyUnlockedRecipes UnlockedRecipesPacketType = 1
	UnlockedRecipesPacketTypeNewlyUnlockedRecipes     UnlockedRecipesPacketType = 2
	UnlockedRecipesPacketTypeRemoveUnlockedRecipes    UnlockedRecipesPacketType = 3
	UnlockedRecipesPacketTypeRemoveAllUnlockedRecipes UnlockedRecipesPacketType = 4
)

type VillageType uint8

const (
	VillageTypeDesert  VillageType = 0
	VillageTypeIce     VillageType = 1
	VillageTypeSavanna VillageType = 2
	VillageTypeTaiga   VillageType = 3
	VillageTypeDefault VillageType = 4
)
