// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientCacheMissResponse struct {
	MissingBlobs []MissingBlobData
}

// Marshal reads or writes ClientCacheMissResponse using its canonical wire layout.
func (x *ClientCacheMissResponse) Marshal(io IO) {
	FuncSlice(io, &x.MissingBlobs, io.Varuint32, func(value *MissingBlobData) {
		value.Marshal(io)
	})
}
