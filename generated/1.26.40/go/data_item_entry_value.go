// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DataItemEntryValue interface {
	isDataItemEntryValue()
}

func marshalDataItemEntryValue(io IO, x *DataItemEntryValue) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				var value DataItemByte
				value.Marshal(io)
				*x = value
			case 1:
				var value DataItemShort
				value.Marshal(io)
				*x = value
			case 2:
				var value DataItemInt
				value.Marshal(io)
				*x = value
			case 3:
				var value DataItemFloat
				value.Marshal(io)
				*x = value
			case 4:
				var value DataItemString
				value.Marshal(io)
				*x = value
			case 5:
				var value DataItemCompoundTag
				value.Marshal(io)
				*x = value
			case 6:
				var value DataItemPos
				value.Marshal(io)
				*x = value
			case 7:
				var value DataItemInt64
				value.Marshal(io)
				*x = value
			case 8:
				var value DataItemVec3
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case DataItemByte:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemShort:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemInt:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemFloat:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemString:
				tag := uint8(4)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemCompoundTag:
				tag := uint8(5)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemPos:
				tag := uint8(6)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemInt64:
				tag := uint8(7)
				io.Uint8(&tag)
				value.Marshal(io)
			case DataItemVec3:
				tag := uint8(8)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
