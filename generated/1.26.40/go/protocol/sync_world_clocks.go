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
				value := new(SyncStateData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(InitializeRegistryData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(AddTimeMarkerData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(RemoveTimeMarkerData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *SyncStateData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *InitializeRegistryData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AddTimeMarkerData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *RemoveTimeMarkerData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
