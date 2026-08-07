// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SimpleEvent struct {
	Type SimpleEventSubtype
}

// Marshal reads or writes SimpleEvent using its canonical wire layout.
func (x *SimpleEvent) Marshal(io IO) {
	enumValue1 := uint16(x.Type)
	io.Uint16(&enumValue1)
	x.Type = SimpleEventSubtype(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
