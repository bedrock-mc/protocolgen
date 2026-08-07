// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CreativeContent struct {
	Groups  []CreativeGroupInfo
	Entries []CreativeItemEntry
}

// Marshal reads or writes CreativeContent using its canonical wire layout.
func (x *CreativeContent) Marshal(io IO) {
	FuncSlice(io, &x.Groups, io.Varuint32, func(value *CreativeGroupInfo) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.Entries, io.Varuint32, func(value *CreativeItemEntry) {
		value.Marshal(io)
	})
}
