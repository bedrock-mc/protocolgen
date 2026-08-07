// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SetScoreInfoItem interface {
	isSetScoreInfoItem()
}

// MarshalSetScoreInfoItem reads or writes the SetScoreInfoItem union using its canonical wire layout.
func MarshalSetScoreInfoItem(io IO, x *SetScoreInfoItem) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(RemoveScore)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ChangePlayerScore)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ChangeEntityScore)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ChangeFakePlayerScore)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *RemoveScore:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangePlayerScore:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangeEntityScore:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangeFakePlayerScore:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
