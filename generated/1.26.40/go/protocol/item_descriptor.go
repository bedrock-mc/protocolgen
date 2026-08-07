// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemDescriptor interface {
	isItemDescriptor()
}

// MarshalItemDescriptor reads or writes the ItemDescriptor union using its canonical wire layout.
func MarshalItemDescriptor(io IO, x *ItemDescriptor) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(InvalidItemDescriptor)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DefaultItemDescriptor)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(MoLangItemDescriptor)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ItemTagItemDescriptor)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *InvalidItemDescriptor:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DefaultItemDescriptor:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *MoLangItemDescriptor:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemTagItemDescriptor:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
