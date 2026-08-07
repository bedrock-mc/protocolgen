// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PropertySyncData struct {
	IntEntriesList   []PropertySyncDataPropertySyncIntEntry
	FloatEntriesList []PropertySyncDataPropertySyncFloatEntry
}

// Marshal reads or writes PropertySyncData using its canonical wire layout.
func (x *PropertySyncData) Marshal(io IO) {
	Slice(io, &x.IntEntriesList)
	Slice(io, &x.FloatEntriesList)
}
