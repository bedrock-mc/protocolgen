// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CreatePhoto struct {
	RawID         uint64
	PhotoName     string
	PhotoItemName string
}

// Marshal reads or writes CreatePhoto using its canonical wire layout.
func (x *CreatePhoto) Marshal(io IO) {
	io.Uint64(&x.RawID)
	io.String(&x.PhotoName)
	io.String(&x.PhotoItemName)
}
