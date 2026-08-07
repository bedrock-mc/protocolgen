// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerListData interface {
	isPlayerListData()
}

// MarshalPlayerListData reads or writes the PlayerListData union using its canonical wire layout.
func MarshalPlayerListData(io IO, x *PlayerListData) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(AddEntry)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(RemoveEntry)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *AddEntry:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *RemoveEntry:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
