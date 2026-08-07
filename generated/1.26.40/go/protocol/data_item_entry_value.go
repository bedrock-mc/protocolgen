// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DataItemEntryValue interface {
	isDataItemEntryValue()
}

// MarshalDataItemEntryValue reads or writes the DataItemEntryValue union using its canonical wire layout.
func MarshalDataItemEntryValue(io IO, x *DataItemEntryValue) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(DataItemByte)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DataItemShort)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(DataItemInt)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(DataItemFloat)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(DataItemString)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(DataItemCompoundTag)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(DataItemPos)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(DataItemInt64)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(DataItemVec3)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *DataItemByte:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemShort:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemInt:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemFloat:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemString:
				tag := uint8(4)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemCompoundTag:
				tag := uint8(5)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemPos:
				tag := uint8(6)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemInt64:
				tag := uint8(7)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemVec3:
				tag := uint8(8)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
