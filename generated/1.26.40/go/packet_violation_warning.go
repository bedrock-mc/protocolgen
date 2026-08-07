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
	enumValue1 := int32(x.ViolationType)
	io.Varint32(&enumValue1)
	x.ViolationType = PacketViolationType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	enumValue2 := int32(x.ViolationSeverity)
	io.Varint32(&enumValue2)
	x.ViolationSeverity = PacketViolationSeverity(enumValue2)
	switch int64(enumValue2) {
	case -1, 0, 1, 2:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	io.Varint32(&x.ViolationPacketId)
	io.String(&x.ViolationContext)
}
