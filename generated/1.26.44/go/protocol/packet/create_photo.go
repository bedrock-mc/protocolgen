// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// CreatePhoto is a packet that allows players to export photos from their portfolios into items in
// their inventory. This packet only works on the Education Edition version of Minecraft.
type CreatePhoto struct {
	RawID uint64
	// PhotoName is the name of the photo.
	PhotoName     string
	PhotoItemName string
}

// Marshal reads or writes CreatePhoto using its canonical wire layout.
func (x *CreatePhoto) Marshal(io protocol.IO) {
	io.Uint64(&x.RawID)
	protocol.Minimum(io, &x.RawID, 0)
	io.String(&x.PhotoName)
	io.String(&x.PhotoItemName)
}

// ID returns the protocol ID for CreatePhoto.
func (*CreatePhoto) ID() uint32 { return IDCreatePhoto }
