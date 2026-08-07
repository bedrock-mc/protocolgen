// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SynchedActorDataCopyableDataList struct {
	Data []DataItemEntry
}

// Marshal reads or writes SynchedActorDataCopyableDataList using its canonical wire layout.
func (x *SynchedActorDataCopyableDataList) Marshal(io IO) {
	Slice(io, &x.Data)
}
