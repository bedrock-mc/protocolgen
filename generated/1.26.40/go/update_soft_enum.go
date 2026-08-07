// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateSoftEnum struct {
	EnumName   string
	Values     []string
	UpdateType SoftEnumUpdateType
}

// Marshal reads or writes UpdateSoftEnum using its canonical wire layout.
func (x *UpdateSoftEnum) Marshal(io IO) {
	io.String(&x.EnumName)
	if !io.Reading() && uint64(len(x.Values)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Values), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Values))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Values = make([]string, int(count1))
	}
	for index2 := range x.Values {
		io.String(&x.Values[index2])
	}
	enumValue3 := uint8(x.UpdateType)
	io.Uint8(&enumValue3)
	x.UpdateType = SoftEnumUpdateType(enumValue3)
	switch int64(enumValue3) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue3, "unknown enum value")
	}
}
