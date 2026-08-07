// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerboundPackSettingChangePackSettingValue interface {
	isServerboundPackSettingChangePackSettingValue()
}

// MarshalServerboundPackSettingChangePackSettingValue reads or writes the ServerboundPackSettingChangePackSettingValue union using its canonical wire layout.
func MarshalServerboundPackSettingChangePackSettingValue(io IO, x *ServerboundPackSettingChangePackSettingValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(ServerboundPackSettingChangePackSettingValueFloat)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ServerboundPackSettingChangePackSettingValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ServerboundPackSettingChangePackSettingValueString)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *ServerboundPackSettingChangePackSettingValueFloat:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ServerboundPackSettingChangePackSettingValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ServerboundPackSettingChangePackSettingValueString:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
