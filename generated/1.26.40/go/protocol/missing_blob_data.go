// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MissingBlobData struct {
	BlobID   uint64
	BlobData []byte
}

// Marshal reads or writes MissingBlobData using its canonical wire layout.
func (x *MissingBlobData) Marshal(io IO) {
	io.Uint64(&x.BlobID)
	io.Bytes(&x.BlobData)
}
