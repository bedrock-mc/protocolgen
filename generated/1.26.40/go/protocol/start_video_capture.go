// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type StartVideoCapture struct {
	FrameRate  uint32
	FilePrefix string
}

func (*StartVideoCapture) isPlayerVideoCaptureData() {}

// Marshal reads or writes StartVideoCapture using its canonical wire layout.
func (x *StartVideoCapture) Marshal(io IO) {
	io.Uint32(&x.FrameRate)
	io.String(&x.FilePrefix)
}
