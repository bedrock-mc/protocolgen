// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BedrockDDUIDataStoreUpdateData interface {
	isBedrockDDUIDataStoreUpdateData()
}

// MarshalBedrockDDUIDataStoreUpdateData reads or writes the BedrockDDUIDataStoreUpdateData union using its canonical wire layout.
func MarshalBedrockDDUIDataStoreUpdateData(io IO, x *BedrockDDUIDataStoreUpdateData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value BedrockDDUIDataStoreUpdateDataDouble
				value.Marshal(io)
				*x = value
			case 1:
				var value BedrockDDUIDataStoreUpdateDataBool
				value.Marshal(io)
				*x = value
			case 2:
				var value BedrockDDUIDataStoreUpdateDataString
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case BedrockDDUIDataStoreUpdateDataDouble:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BedrockDDUIDataStoreUpdateDataBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BedrockDDUIDataStoreUpdateDataString:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
