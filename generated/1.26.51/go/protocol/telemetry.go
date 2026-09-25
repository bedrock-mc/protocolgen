// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacyTelemetryType int32

const (
	LegacyTelemetryTypeAchievement                     LegacyTelemetryType = 0
	LegacyTelemetryTypeInteraction                     LegacyTelemetryType = 1
	LegacyTelemetryTypePortalcreated                   LegacyTelemetryType = 2
	LegacyTelemetryTypePortalused                      LegacyTelemetryType = 3
	LegacyTelemetryTypeMobkilled                       LegacyTelemetryType = 4
	LegacyTelemetryTypeCauldronused                    LegacyTelemetryType = 5
	LegacyTelemetryTypePlayerdied                      LegacyTelemetryType = 6
	LegacyTelemetryTypeBosskilled                      LegacyTelemetryType = 7
	LegacyTelemetryTypeAgentcommandObsolete            LegacyTelemetryType = 8
	LegacyTelemetryTypeAgentcreated                    LegacyTelemetryType = 9
	LegacyTelemetryTypePatternremovedObsolete          LegacyTelemetryType = 10
	LegacyTelemetryTypeSlashcommand                    LegacyTelemetryType = 11
	LegacyTelemetryTypeFishbucketedObsolete            LegacyTelemetryType = 12
	LegacyTelemetryTypeMobborn                         LegacyTelemetryType = 13
	LegacyTelemetryTypePetdiedObsolete                 LegacyTelemetryType = 14
	LegacyTelemetryTypePoicauldronused                 LegacyTelemetryType = 15
	LegacyTelemetryTypeComposterused                   LegacyTelemetryType = 16
	LegacyTelemetryTypeBellused                        LegacyTelemetryType = 17
	LegacyTelemetryTypeActordefinition                 LegacyTelemetryType = 18
	LegacyTelemetryTypeRaidupdate                      LegacyTelemetryType = 19
	LegacyTelemetryTypePlayermovementanomalyObsolete   LegacyTelemetryType = 20
	LegacyTelemetryTypePlayermovementcorrectedObsolete LegacyTelemetryType = 21
	LegacyTelemetryTypeHoneyharvested                  LegacyTelemetryType = 22
	LegacyTelemetryTypeTargetblockhit                  LegacyTelemetryType = 23
	LegacyTelemetryTypePiglinbarter                    LegacyTelemetryType = 24
	LegacyTelemetryTypePlayerwaxedorunwaxedcopper      LegacyTelemetryType = 25
	LegacyTelemetryTypeCodebuilderruntimeaction        LegacyTelemetryType = 26
	LegacyTelemetryTypeCodebuilderscoreboard           LegacyTelemetryType = 27
	LegacyTelemetryTypeStriderriddeninlavainoverworld  LegacyTelemetryType = 28
	LegacyTelemetryTypeSneakclosetosculksensor         LegacyTelemetryType = 29
	LegacyTelemetryTypeCarefulrestoration              LegacyTelemetryType = 30
	LegacyTelemetryTypeItemused                        LegacyTelemetryType = 31
)

// Marshal reads or writes LegacyTelemetryType through its int32 wire encoding.
func (x *LegacyTelemetryType) Marshal(io IO) { io.Varint32((*int32)(x)) }
