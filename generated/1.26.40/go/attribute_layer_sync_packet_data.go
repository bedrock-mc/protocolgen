// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AttributeLayerSyncPacketData interface {
	isAttributeLayerSyncPacketData()
}

func marshalAttributeLayerSyncPacketData(io IO, x *AttributeLayerSyncPacketData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value AttributeLayerSyncPacketDataUpdateAttributeLayersData
				value.Marshal(io)
				*x = value
			case 1:
				var value AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData
				value.Marshal(io)
				*x = value
			case 2:
				var value AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData
				value.Marshal(io)
				*x = value
			case 3:
				var value AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case AttributeLayerSyncPacketDataUpdateAttributeLayersData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
