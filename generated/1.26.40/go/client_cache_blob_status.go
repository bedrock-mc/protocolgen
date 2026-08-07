// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientCacheBlobStatus struct {
	MissingIds []uint64
	FoundIds   []uint64
}

// Marshal reads or writes ClientCacheBlobStatus using its canonical wire layout.
func (x *ClientCacheBlobStatus) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.MissingIds)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MissingIds), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.MissingIds))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.MissingIds = make([]uint64, int(count1))
	}
	for index2 := range x.MissingIds {
		io.Uint64(&x.MissingIds[index2])
	}
	if !io.Reading() && uint64(len(x.FoundIds)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.FoundIds), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.FoundIds))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.FoundIds = make([]uint64, int(count3))
	}
	for index4 := range x.FoundIds {
		io.Uint64(&x.FoundIds[index4])
	}
}
