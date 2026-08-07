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
	enumValue1 := uint8(x.Type)
	io.Uint8(&enumValue1)
	x.Type = PhotoType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	enumValue2 := uint8(x.SourceType)
	io.Uint8(&enumValue2)
	x.SourceType = PhotoType(enumValue2)
	switch int64(enumValue2) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	io.Int64(&x.OwnerID)
	io.String(&x.NewPhotoName)
}
