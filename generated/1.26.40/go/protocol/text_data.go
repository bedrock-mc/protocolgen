// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TextData interface {
	isTextData()
}

// MarshalTextData reads or writes the TextData union using its canonical wire layout.
func MarshalTextData(io IO, x *TextData) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(MessageOnly)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(AuthorAndMessage)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(MessageAndParams)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(TextDataPopup)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(TextDataJukeboxPopup)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(TextDataTip)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(TextDataSystemMessage)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(TextDataWhisper)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(TextDataAnnouncement)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(TextDataTextObjectWhisper)
				value.Marshal(io)
				*x = value
			case 10:
				value := new(TextDataTextObject)
				value.Marshal(io)
				*x = value
			case 11:
				value := new(TextDataTextObjectAnnouncement)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *MessageOnly:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *AuthorAndMessage:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *MessageAndParams:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataPopup:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataJukeboxPopup:
				tag := uint8(4)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataTip:
				tag := uint8(5)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataSystemMessage:
				tag := uint8(6)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataWhisper:
				tag := uint8(7)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataAnnouncement:
				tag := uint8(8)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataTextObjectWhisper:
				tag := uint8(9)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataTextObject:
				tag := uint8(10)
				io.Uint8(&tag)
				value.Marshal(io)
			case *TextDataTextObjectAnnouncement:
				tag := uint8(11)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
