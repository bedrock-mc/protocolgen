// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

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
