// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncWorldClocks struct {
	Data SyncWorldClocksData
}

// Marshal reads or writes SyncWorldClocks using its canonical wire layout.
func (x *SyncWorldClocks) Marshal(io IO) {
	marshalSyncWorldClocksData(io, &x.Data)
}
