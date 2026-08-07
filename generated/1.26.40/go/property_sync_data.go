// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PropertySyncData struct {
	IntEntriesList   []PropertySyncDataPropertySyncIntEntry
	FloatEntriesList []PropertySyncDataPropertySyncFloatEntry
}

// Marshal reads or writes PropertySyncData using its canonical wire layout.
func (x *PropertySyncData) Marshal(io IO) {
	FuncSlice(io, &x.IntEntriesList, io.Varuint32, func(value *PropertySyncDataPropertySyncIntEntry) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.FloatEntriesList, io.Varuint32, func(value *PropertySyncDataPropertySyncFloatEntry) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
