// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextBody interface {
	isTextBody()
}

// MarshalTextBody reads or writes the TextBody union using its canonical wire layout.
func MarshalTextBody(io IO, x *TextBody) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				var value TextMessageOnly
				value.Marshal(io)
				*x = value
			case 1:
				var value TextAuthorAndMessage
				value.Marshal(io)
				*x = value
			case 2:
				var value TextMessageAndParams
				value.Marshal(io)
				*x = value
			case 3:
				var value TextBodyPopup
				value.Marshal(io)
				*x = value
			case 4:
				var value TextBodyJukeboxPopup
				value.Marshal(io)
				*x = value
			case 5:
				var value TextBodyTip
				value.Marshal(io)
				*x = value
			case 6:
				var value TextBodySystemMessage
				value.Marshal(io)
				*x = value
			case 7:
				var value TextBodyWhisper
				value.Marshal(io)
				*x = value
			case 8:
				var value TextBodyAnnouncement
				value.Marshal(io)
				*x = value
			case 9:
				var value TextBodyTextObjectWhisper
				value.Marshal(io)
				*x = value
			case 10:
				var value TextBodyTextObject
				value.Marshal(io)
				*x = value
			case 11:
				var value TextBodyTextObjectAnnouncement
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case TextMessageOnly:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextAuthorAndMessage:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextMessageAndParams:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyPopup:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyJukeboxPopup:
				tag := uint8(4)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyTip:
				tag := uint8(5)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodySystemMessage:
				tag := uint8(6)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyWhisper:
				tag := uint8(7)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyAnnouncement:
				tag := uint8(8)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyTextObjectWhisper:
				tag := uint8(9)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyTextObject:
				tag := uint8(10)
				io.Uint8(&tag)
				value.Marshal(io)
			case TextBodyTextObjectAnnouncement:
				tag := uint8(11)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
