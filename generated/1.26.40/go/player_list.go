// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerList struct {
	Entries []PlayerListEntriesItem
}

// Marshal reads or writes PlayerList using its canonical wire layout.
func (x *PlayerList) Marshal(io IO) {
	FuncSlice(io, &x.Entries, io.Varuint32, func(value *PlayerListEntriesItem) {
		item := *value
		marshalPlayerListEntriesItem(io, &item)
		*value = item
	})
}
