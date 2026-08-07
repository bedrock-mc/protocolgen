// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerLocationData interface {
	isPlayerLocationData()
}

// MarshalPlayerLocationData reads or writes the PlayerLocationData union using its canonical wire layout.
func MarshalPlayerLocationData(io IO, x *PlayerLocationData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(CoordinatesLocation)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(HiddenLocation)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *CoordinatesLocation:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *HiddenLocation:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
