// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PrimitiveShapeDataExtraShapeData interface {
	isPrimitiveShapeDataExtraShapeData()
}

func marshalPrimitiveShapeDataExtraShapeData(io IO, x *PrimitiveShapeDataExtraShapeData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value PrimitiveShapeDataExtraShapeDataEmpty0
				value.Marshal(io)
				*x = value
			case 1:
				var value ArrowData
				value.Marshal(io)
				*x = value
			case 2:
				var value TextData
				value.Marshal(io)
				*x = value
			case 3:
				var value BoxData
				value.Marshal(io)
				*x = value
			case 4:
				var value LineData
				value.Marshal(io)
				*x = value
			case 5:
				var value SphereData
				value.Marshal(io)
				*x = value
			case 6:
				var value CylinderData
				value.Marshal(io)
				*x = value
			case 7:
				var value PyramidData
				value.Marshal(io)
				*x = value
			case 8:
				var value EllipsoidData
				value.Marshal(io)
				*x = value
			case 9:
				var value ConeData
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case PrimitiveShapeDataExtraShapeDataEmpty0:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ArrowData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case TextData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BoxData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case LineData:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SphereData:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case CylinderData:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case PyramidData:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case EllipsoidData:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ConeData:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
