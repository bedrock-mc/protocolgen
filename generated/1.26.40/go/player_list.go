// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerList struct {
	Entries []PlayerListEntriesItem
}

// Marshal reads or writes PlayerList using its canonical wire layout.
func (x *PlayerList) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Entries)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Entries), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Entries))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Entries = make([]PlayerListEntriesItem, int(count1))
	}
	for index2 := range x.Entries {
		marshalPlayerListEntriesItem(io, &x.Entries[index2])
	}
}
