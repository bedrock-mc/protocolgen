// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerVideoCapture struct {
	Action protocol.PlayerVideoCaptureAction
}

// Marshal reads or writes PlayerVideoCapture using its canonical wire layout.
func (x *PlayerVideoCapture) Marshal(io protocol.IO) {
	protocol.MarshalPlayerVideoCaptureAction(io, &x.Action)
}

// ID returns the protocol ID for PlayerVideoCapture.
func (*PlayerVideoCapture) ID() uint32 { return IDPlayerVideoCapture }
