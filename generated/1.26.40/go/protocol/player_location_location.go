// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerLocationLocation interface {
	isPlayerLocationLocation()
}

// MarshalPlayerLocationLocation reads or writes the PlayerLocationLocation union using its canonical wire layout.
func MarshalPlayerLocationLocation(io IO, x *PlayerLocationLocation) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(PlayerLocationCoordinatesLocation)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(PlayerLocationHiddenLocation)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *PlayerLocationCoordinatesLocation:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PlayerLocationHiddenLocation:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
