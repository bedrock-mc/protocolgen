// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PacketViolationWarning struct {
	ViolationType     PacketViolationType
	ViolationSeverity PacketViolationSeverity
	ViolationPacketId int32
	ViolationContext  string
}

// Marshal reads or writes PacketViolationWarning using its canonical wire layout.
func (x *PacketViolationWarning) Marshal(io IO) {
	IntegerFunc(&x.ViolationType, io.Varint32)
	IntegerFunc(&x.ViolationSeverity, io.Varint32)
	io.Varint32(&x.ViolationPacketId)
	io.String(&x.ViolationContext)
}
