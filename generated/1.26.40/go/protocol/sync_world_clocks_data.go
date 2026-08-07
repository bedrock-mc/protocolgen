// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncWorldClocksData interface {
	isSyncWorldClocksData()
}

// MarshalSyncWorldClocksData reads or writes the SyncWorldClocksData union using its canonical wire layout.
func MarshalSyncWorldClocksData(io IO, x *SyncWorldClocksData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value SyncWorldClocksSyncStateData
				value.Marshal(io)
				*x = value
			case 1:
				var value SyncWorldClocksInitializeRegistryData
				value.Marshal(io)
				*x = value
			case 2:
				var value SyncWorldClocksAddTimeMarkerData
				value.Marshal(io)
				*x = value
			case 3:
				var value SyncWorldClocksRemoveTimeMarkerData
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case SyncWorldClocksSyncStateData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SyncWorldClocksInitializeRegistryData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SyncWorldClocksAddTimeMarkerData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SyncWorldClocksRemoveTimeMarkerData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
