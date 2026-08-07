// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerVideoCaptureAction interface {
	isPlayerVideoCaptureAction()
}

// MarshalPlayerVideoCaptureAction reads or writes the PlayerVideoCaptureAction union using its canonical wire layout.
func MarshalPlayerVideoCaptureAction(io IO, x *PlayerVideoCaptureAction) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(PlayerVideoCaptureStopVideoCapture)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(PlayerVideoCaptureStartVideoCapture)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *PlayerVideoCaptureStopVideoCapture:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *PlayerVideoCaptureStartVideoCapture:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
