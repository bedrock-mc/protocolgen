// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerUpdateEntityOverridesData interface {
	isPlayerUpdateEntityOverridesData()
}

// MarshalPlayerUpdateEntityOverridesData reads or writes the PlayerUpdateEntityOverridesData union using its canonical wire layout.
func MarshalPlayerUpdateEntityOverridesData(io IO, x *PlayerUpdateEntityOverridesData) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(ClearOverride)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(RemoveOverride)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(IntOverride)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(FloatOverride)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *ClearOverride:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *RemoveOverride:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *IntOverride:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *FloatOverride:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
