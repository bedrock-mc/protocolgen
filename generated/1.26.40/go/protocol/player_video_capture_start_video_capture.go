// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerVideoCaptureStartVideoCapture struct {
	FrameRate  uint32
	FilePrefix string
}

func (*PlayerVideoCaptureStartVideoCapture) isPlayerVideoCaptureAction() {}

// Marshal reads or writes PlayerVideoCaptureStartVideoCapture using its canonical wire layout.
func (x *PlayerVideoCaptureStartVideoCapture) Marshal(io IO) {
	io.Uint32(&x.FrameRate)
	io.String(&x.FilePrefix)
}
