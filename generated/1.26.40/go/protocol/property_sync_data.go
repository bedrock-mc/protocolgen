// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PropertySyncData struct {
	IntEntriesList   []PropertySyncDataPropertySyncIntEntry
	FloatEntriesList []PropertySyncDataPropertySyncFloatEntry
}

// Marshal reads or writes PropertySyncData using its canonical wire layout.
func (x *PropertySyncData) Marshal(io IO) {
	FuncSlice(io, &x.IntEntriesList, io.Varuint32, func(value *PropertySyncDataPropertySyncIntEntry) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.FloatEntriesList, io.Varuint32, func(value *PropertySyncDataPropertySyncFloatEntry) {
		value.Marshal(io)
	})
}
