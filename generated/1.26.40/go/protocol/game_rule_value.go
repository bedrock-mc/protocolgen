// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleValue interface {
	isGameRuleValue()
}

// MarshalGameRuleValue reads or writes the GameRuleValue union using its canonical wire layout.
func MarshalGameRuleValue(io IO, x *GameRuleValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(GameRuleValueEmpty)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(GameRuleValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(GameRuleValueInt32)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(GameRuleValueFloat)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *GameRuleValueEmpty:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueInt32:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueFloat:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
