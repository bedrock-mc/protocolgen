// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BedrockDDUI interface {
	isBedrockDDUI()
}

func marshalBedrockDDUI(io IO, x *BedrockDDUI) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value BedrockDDUIDataStoreUpdate
				value.Marshal(io)
				*x = value
			case 1:
				var value BedrockDDUIDataStoreChange
				value.Marshal(io)
				*x = value
			case 2:
				var value BedrockDDUIDataStoreRemoval
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case BedrockDDUIDataStoreUpdate:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BedrockDDUIDataStoreChange:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BedrockDDUIDataStoreRemoval:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
