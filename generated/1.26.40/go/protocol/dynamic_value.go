// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValue interface {
	isDynamicValue()
}

// MarshalDynamicValue reads or writes the DynamicValue union using its canonical wire layout.
func MarshalDynamicValue(io IO, x *DynamicValue) {
	UnionFunc(io,
		func() {
			var tag int32
			io.Int32(&tag)
			switch int64(tag) {
			case 0:
				value := new(DynamicValueNone)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DynamicValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(DynamicValueInt64)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(DynamicValueDouble)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(DynamicValueString)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(DynamicValueList)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(DynamicValueMap)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *DynamicValueNone:
				tag := int32(0)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueBool:
				tag := int32(1)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueInt64:
				tag := int32(2)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueDouble:
				tag := int32(3)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueString:
				tag := int32(4)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueList:
				tag := int32(5)
				io.Int32(&tag)
				value.Marshal(io)
			case *DynamicValueMap:
				tag := int32(6)
				io.Int32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
