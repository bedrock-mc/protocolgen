// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValue interface {
	isCerealDynamicValue()
}

// MarshalCerealDynamicValue reads or writes the CerealDynamicValue union using its canonical wire layout.
func MarshalCerealDynamicValue(io IO, x *CerealDynamicValue) {
	UnionFunc(io,
		func() {
			var tag int32
			io.Int32(&tag)
			switch int64(tag) {
			case 0:
				value := new(CerealDynamicValueNone)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(CerealDynamicValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(CerealDynamicValueInt64)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(CerealDynamicValueDouble)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(CerealDynamicValueString)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(CerealDynamicValueList)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(CerealDynamicValueMap)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *CerealDynamicValueNone:
				tag := int32(0)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueBool:
				tag := int32(1)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueInt64:
				tag := int32(2)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueDouble:
				tag := int32(3)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueString:
				tag := int32(4)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueList:
				tag := int32(5)
				io.Int32(&tag)
				value.Marshal(io)
			case *CerealDynamicValueMap:
				tag := int32(6)
				io.Int32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
