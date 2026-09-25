// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextData interface {
	Marshaler
	tagTextData() uint32
}

// MarshalTextData reads or writes the TextData union using its canonical wire layout.
func MarshalTextData(io IO, x *TextData) {
	Union(io, x, io.Varuint32, TextData.tagTextData, func(tag uint32) TextData {
		switch tag {
		case 0:
			return new(MessageOnly)
		case 1:
			return new(AuthorAndMessage)
		case 2:
			return new(MessageAndParams)
		}
		return nil
	})
}

type TextPacketType uint8

const (
	TextPacketTypeRaw                    TextPacketType = 0
	TextPacketTypeTip                    TextPacketType = 5
	TextPacketTypeSystemmessage          TextPacketType = 6
	TextPacketTypeTextobjectwhisper      TextPacketType = 9
	TextPacketTypeTextobject             TextPacketType = 10
	TextPacketTypeTextobjectannouncement TextPacketType = 11
)

// Marshal reads or writes TextPacketType through its uint8 wire encoding.
func (x *TextPacketType) Marshal(io IO) { io.Uint8((*uint8)(x)) }

type TextProcessingEventOrigin int32

const (
	TextProcessingEventOriginUnknown            TextProcessingEventOrigin = -1
	TextProcessingEventOriginServerchatpublic   TextProcessingEventOrigin = 0
	TextProcessingEventOriginServerchatwhisper  TextProcessingEventOrigin = 1
	TextProcessingEventOriginSigntext           TextProcessingEventOrigin = 2
	TextProcessingEventOriginAnviltext          TextProcessingEventOrigin = 3
	TextProcessingEventOriginBookandquilltext   TextProcessingEventOrigin = 4
	TextProcessingEventOriginCommandblocktext   TextProcessingEventOrigin = 5
	TextProcessingEventOriginBlockactordatatext TextProcessingEventOrigin = 6
	TextProcessingEventOriginJoineventtext      TextProcessingEventOrigin = 7
	TextProcessingEventOriginLeaveeventtext     TextProcessingEventOrigin = 8
	TextProcessingEventOriginSlashcommandchat   TextProcessingEventOrigin = 9
	TextProcessingEventOriginCartographytext    TextProcessingEventOrigin = 10
	TextProcessingEventOriginKickcommand        TextProcessingEventOrigin = 11
	TextProcessingEventOriginTitlecommand       TextProcessingEventOrigin = 12
	TextProcessingEventOriginSummoncommand      TextProcessingEventOrigin = 13
	TextProcessingEventOriginServerform         TextProcessingEventOrigin = 14
	TextProcessingEventOriginDatadrivenui       TextProcessingEventOrigin = 15
)

// Marshal reads or writes TextProcessingEventOrigin through its int32 wire encoding.
func (x *TextProcessingEventOrigin) Marshal(io IO) { io.Int32((*int32)(x)) }
