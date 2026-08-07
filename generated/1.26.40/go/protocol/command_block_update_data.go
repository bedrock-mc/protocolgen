// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandBlockUpdateData interface {
	isCommandBlockUpdateData()
}

// MarshalCommandBlockUpdateData reads or writes the CommandBlockUpdateData union using its canonical wire layout.
func MarshalCommandBlockUpdateData(io IO, x *CommandBlockUpdateData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(EntityCommandTarget)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(BlockCommandData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *EntityCommandTarget:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BlockCommandData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
