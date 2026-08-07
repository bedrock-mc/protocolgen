// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientCacheBlobStatus struct {
	MissingIds []uint64
	FoundIds   []uint64
}

// Marshal reads or writes ClientCacheBlobStatus using its canonical wire layout.
func (x *ClientCacheBlobStatus) Marshal(io IO) {
	FuncSlice(io, &x.MissingIds, io.Varuint32, io.Uint64)
	FuncSlice(io, &x.FoundIds, io.Varuint32, io.Uint64)
}
