// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerUpdateEntityOverridesUpdate interface {
	isPlayerUpdateEntityOverridesUpdate()
}

// MarshalPlayerUpdateEntityOverridesUpdate reads or writes the PlayerUpdateEntityOverridesUpdate union using its canonical wire layout.
func MarshalPlayerUpdateEntityOverridesUpdate(io IO, x *PlayerUpdateEntityOverridesUpdate) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				var value PlayerUpdateEntityOverridesClearOverride
				value.Marshal(io)
				*x = value
			case 1:
				var value PlayerUpdateEntityOverridesRemoveOverride
				value.Marshal(io)
				*x = value
			case 2:
				var value PlayerUpdateEntityOverridesIntOverride
				value.Marshal(io)
				*x = value
			case 3:
				var value PlayerUpdateEntityOverridesFloatOverride
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case PlayerUpdateEntityOverridesClearOverride:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case PlayerUpdateEntityOverridesRemoveOverride:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case PlayerUpdateEntityOverridesIntOverride:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case PlayerUpdateEntityOverridesFloatOverride:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
