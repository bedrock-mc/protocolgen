// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CreativeContent struct {
	Groups  []CreativeGroupInfo
	Entries []CreativeItemEntry
}

// Marshal reads or writes CreativeContent using its canonical wire layout.
func (x *CreativeContent) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Groups)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Groups), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Groups))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Groups = make([]CreativeGroupInfo, int(count1))
	}
	for index2 := range x.Groups {
		x.Groups[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Entries)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Entries), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.Entries))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.Entries = make([]CreativeItemEntry, int(count3))
	}
	for index4 := range x.Entries {
		x.Entries[index4].Marshal(io)
	}
}
