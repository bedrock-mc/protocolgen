// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleRuleValue interface {
	isGameRuleRuleValue()
}

// MarshalGameRuleRuleValue reads or writes the GameRuleRuleValue union using its canonical wire layout.
func MarshalGameRuleRuleValue(io IO, x *GameRuleRuleValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(GameRuleRuleValueEmpty0)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(GameRuleRuleValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(GameRuleRuleValueInt32)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(GameRuleRuleValueFloat)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *GameRuleRuleValueEmpty0:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleRuleValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleRuleValueInt32:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleRuleValueFloat:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
