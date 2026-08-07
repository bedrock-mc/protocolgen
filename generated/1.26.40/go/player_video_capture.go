// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerVideoCapture struct {
	Action PlayerVideoCaptureAction
}

// Marshal reads or writes PlayerVideoCapture using its canonical wire layout.
func (x *PlayerVideoCapture) Marshal(io IO) {
	marshalPlayerVideoCaptureAction(io, &x.Action)
}
