// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetHud struct {
	HudElement []HudElement
	HudVisible HudVisibility
}

// Marshal reads or writes SetHud using its canonical wire layout.
func (x *SetHud) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.HudElement)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.HudElement), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.HudElement))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.HudElement = make([]HudElement, int(count1))
	}
	for index2 := range x.HudElement {
		enumValue3 := int32(x.HudElement[index2])
		io.Varint32(&enumValue3)
		x.HudElement[index2] = HudElement(enumValue3)
		switch int64(enumValue3) {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		default:
			io.InvalidValue(enumValue3, "unknown enum value")
		}
	}
	enumValue4 := int32(x.HudVisible)
	io.Varint32(&enumValue4)
	x.HudVisible = HudVisibility(enumValue4)
	switch int64(enumValue4) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue4, "unknown enum value")
	}
}
