// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MaterialReducerDataEntry struct {
	FromItemKey      int32
	ItemIdsAndCounts []MaterialReducerEntryOutput
}

// Marshal reads or writes MaterialReducerDataEntry using its canonical wire layout.
func (x *MaterialReducerDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemKey)
	FuncSlice(io, &x.ItemIdsAndCounts, io.Varuint32, func(value *MaterialReducerEntryOutput) {
		value.Marshal(io)
	})
}
