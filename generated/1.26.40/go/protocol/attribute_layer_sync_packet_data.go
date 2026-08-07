// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeLayerSyncPacketData interface {
	isAttributeLayerSyncPacketData()
}

// MarshalAttributeLayerSyncPacketData reads or writes the AttributeLayerSyncPacketData union using its canonical wire layout.
func MarshalAttributeLayerSyncPacketData(io IO, x *AttributeLayerSyncPacketData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(AttributeLayerSyncPacketDataUpdateAttributeLayersData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *AttributeLayerSyncPacketDataUpdateAttributeLayersData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
