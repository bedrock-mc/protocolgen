// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CreatePhoto struct {
	RawID         uint64
	PhotoName     string
	PhotoItemName string
}

// Marshal reads or writes CreatePhoto using its canonical wire layout.
func (x *CreatePhoto) Marshal(io protocol.IO) {
	io.Uint64(&x.RawID)
	io.String(&x.PhotoName)
	io.String(&x.PhotoItemName)
}

// ID returns the protocol ID for CreatePhoto.
func (*CreatePhoto) ID() uint32 { return IDCreatePhoto }
