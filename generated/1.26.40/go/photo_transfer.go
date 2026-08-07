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
