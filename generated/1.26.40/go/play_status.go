// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayStatus struct {
	Status PlayStatusType
}

// Marshal reads or writes PlayStatus using its canonical wire layout.
func (x *PlayStatus) Marshal(io IO) {
	enumValue1 := int32(x.Status)
	io.BEInt32(&enumValue1)
	x.Status = PlayStatusType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
