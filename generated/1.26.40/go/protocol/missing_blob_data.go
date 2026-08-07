// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MissingBlobData struct {
	BlobId   uint64
	BlobData []byte
}

// Marshal reads or writes MissingBlobData using its canonical wire layout.
func (x *MissingBlobData) Marshal(io IO) {
	io.Uint64(&x.BlobId)
	io.Bytes(&x.BlobData)
}
