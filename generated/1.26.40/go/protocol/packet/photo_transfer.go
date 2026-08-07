// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PhotoTransfer struct {
	PhotoName    string
	PhotoData    []byte
	BookID       string
	Type         protocol.PhotoType
	SourceType   protocol.PhotoType
	OwnerID      int64
	NewPhotoName string
}

// Marshal reads or writes PhotoTransfer using its canonical wire layout.
func (x *PhotoTransfer) Marshal(io protocol.IO) {
	io.String(&x.PhotoName)
	io.Bytes(&x.PhotoData)
	io.String(&x.BookID)
	protocol.IntegerFunc(&x.Type, io.Uint8)
	protocol.IntegerFunc(&x.SourceType, io.Uint8)
	io.Int64(&x.OwnerID)
	io.String(&x.NewPhotoName)
}

// ID returns the protocol ID for PhotoTransfer.
func (*PhotoTransfer) ID() uint32 { return IDPhotoTransfer }
