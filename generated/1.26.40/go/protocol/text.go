// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextData interface {
	Marshaler
	tagTextData() uint8
}

// MarshalTextData reads or writes the TextData union using its canonical wire layout.
func MarshalTextData(io IO, x *TextData) {
	Union(io, x, io.Uint8, TextData.tagTextData, func(tag uint8) TextData {
		switch tag {
		case 0:
			return new(MessageOnly)
		case 1:
			return new(AuthorAndMessage)
		case 2:
			return new(MessageAndParams)
		case 3:
			return new(TextDataPopup)
		case 4:
			return new(TextDataJukeboxPopup)
		case 5:
			return new(TextDataTip)
		case 6:
			return new(TextDataSystemMessage)
		case 7:
			return new(TextDataWhisper)
		case 8:
			return new(TextDataAnnouncement)
		case 9:
			return new(TextDataTextObjectWhisper)
		case 10:
			return new(TextDataTextObject)
		case 11:
			return new(TextDataTextObjectAnnouncement)
		}
		return nil
	})
}

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

// Marshal reads or writes TextProcessingEventOrigin through its int32 wire encoding.
func (x *TextProcessingEventOrigin) Marshal(io IO) { io.Int32((*int32)(x)) }
