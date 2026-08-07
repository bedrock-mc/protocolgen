// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerVideoCaptureData interface {
	isPlayerVideoCaptureData()
}

// MarshalPlayerVideoCaptureData reads or writes the PlayerVideoCaptureData union using its canonical wire layout.
func MarshalPlayerVideoCaptureData(io IO, x *PlayerVideoCaptureData) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(StopVideoCapture)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(StartVideoCapture)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *StopVideoCapture:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *StartVideoCapture:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
