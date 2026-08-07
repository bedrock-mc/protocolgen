// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DisconnectMessages interface {
	isDisconnectMessages()
}

// MarshalDisconnectMessages reads or writes the DisconnectMessages union using its canonical wire layout.
func MarshalDisconnectMessages(io IO, x *DisconnectMessages) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value DisconnectPacketMessages
				value.Marshal(io)
				*x = value
			case 1:
				var value DisconnectMessagesEmpty1
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case DisconnectPacketMessages:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case DisconnectMessagesEmpty1:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
