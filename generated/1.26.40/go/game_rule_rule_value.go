// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GameRuleRuleValue interface {
	isGameRuleRuleValue()
}

func marshalGameRuleRuleValue(io IO, x *GameRuleRuleValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value GameRuleRuleValueEmpty0
				value.Marshal(io)
				*x = value
			case 1:
				var value GameRuleRuleValueBool
				value.Marshal(io)
				*x = value
			case 2:
				var value GameRuleRuleValueInt32
				value.Marshal(io)
				*x = value
			case 3:
				var value GameRuleRuleValueFloat
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case GameRuleRuleValueEmpty0:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case GameRuleRuleValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case GameRuleRuleValueInt32:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case GameRuleRuleValueFloat:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
