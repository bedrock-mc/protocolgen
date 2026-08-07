// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetScoreScoreInfoItem interface {
	isSetScoreScoreInfoItem()
}

func marshalSetScoreScoreInfoItem(io IO, x *SetScoreScoreInfoItem) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				var value RemoveScore
				value.Marshal(io)
				*x = value
			case 1:
				var value ChangePlayerScore
				value.Marshal(io)
				*x = value
			case 2:
				var value ChangeEntityScore
				value.Marshal(io)
				*x = value
			case 3:
				var value ChangeFakePlayerScore
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case RemoveScore:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case ChangePlayerScore:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case ChangeEntityScore:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case ChangeFakePlayerScore:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
