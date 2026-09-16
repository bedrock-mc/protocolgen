// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncWorldClocksData interface {
	Marshaler
	tagSyncWorldClocksData() uint32
}

// MarshalSyncWorldClocksData reads or writes the SyncWorldClocksData union using its canonical wire layout.
func MarshalSyncWorldClocksData(io IO, x *SyncWorldClocksData) {
	Union(io, x, io.Varuint32, SyncWorldClocksData.tagSyncWorldClocksData, func(tag uint32) SyncWorldClocksData {
		switch tag {
		case 0:
			return new(SyncStateData)
		case 1:
			return new(InitializeRegistryData)
		case 2:
			return new(AddTimeMarkerData)
		case 3:
			return new(RemoveTimeMarkerData)
		}
		return nil
	})
}
