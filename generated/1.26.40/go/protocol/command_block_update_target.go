// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandBlockUpdateTarget interface {
	isCommandBlockUpdateTarget()
}

// MarshalCommandBlockUpdateTarget reads or writes the CommandBlockUpdateTarget union using its canonical wire layout.
func MarshalCommandBlockUpdateTarget(io IO, x *CommandBlockUpdateTarget) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value CommandBlockUpdateEntityCommandTarget
				value.Marshal(io)
				*x = value
			case 1:
				var value CommandBlockUpdateBlockCommandData
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case CommandBlockUpdateEntityCommandTarget:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case CommandBlockUpdateBlockCommandData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
