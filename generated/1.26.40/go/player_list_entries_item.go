// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerListEntriesItem interface {
	isPlayerListEntriesItem()
}

func marshalPlayerListEntriesItem(io IO, x *PlayerListEntriesItem) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				var value PlayerListAddEntry
				value.Marshal(io)
				*x = value
			case 1:
				var value PlayerListRemoveEntry
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case PlayerListAddEntry:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case PlayerListRemoveEntry:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
