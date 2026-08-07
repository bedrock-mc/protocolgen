// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PhotoTransfer struct {
	PhotoName    string
	PhotoData    []byte
	BookID       string
	Type         PhotoType
	SourceType   PhotoType
	OwnerID      int64
	NewPhotoName string
}

// Marshal reads or writes PhotoTransfer using its canonical wire layout.
func (x *PhotoTransfer) Marshal(io IO) {
	io.String(&x.PhotoName)
	io.Bytes(&x.PhotoData)
	io.String(&x.BookID)
	IntegerFunc(&x.Type, io.Uint8)
	IntegerFunc(&x.SourceType, io.Uint8)
	io.Int64(&x.OwnerID)
	io.String(&x.NewPhotoName)
}
