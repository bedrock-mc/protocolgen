// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PrimitiveShapeDataExtraShapeData interface {
	isPrimitiveShapeDataExtraShapeData()
}

// MarshalPrimitiveShapeDataExtraShapeData reads or writes the PrimitiveShapeDataExtraShapeData union using its canonical wire layout.
func MarshalPrimitiveShapeDataExtraShapeData(io IO, x *PrimitiveShapeDataExtraShapeData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(PrimitiveShapeDataExtraShapeDataEmpty0)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ArrowData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(TextData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(BoxData)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(LineData)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(SphereData)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(CylinderData)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(PyramidData)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(EllipsoidData)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(ConeData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *PrimitiveShapeDataExtraShapeDataEmpty0:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ArrowData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *TextData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BoxData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *LineData:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SphereData:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CylinderData:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PyramidData:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *EllipsoidData:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ConeData:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
