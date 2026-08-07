// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerSyncData interface {
	isAttributeLayerSyncData()
}

// MarshalAttributeLayerSyncData reads or writes the AttributeLayerSyncData union using its canonical wire layout.
func MarshalAttributeLayerSyncData(io IO, x *AttributeLayerSyncData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(AttributeLayerData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(AttributeLayerSettings)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(EnvironmentAttributeData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(RemoveEnvironmentAttributes)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *AttributeLayerData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AttributeLayerSettings:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *EnvironmentAttributeData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *RemoveEnvironmentAttributes:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
