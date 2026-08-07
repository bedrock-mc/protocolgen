// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SimulationType struct {
	SimType SimulationTypeType
}

// Marshal reads or writes SimulationType using its canonical wire layout.
func (x *SimulationType) Marshal(io IO) {
	enumValue1 := uint8(x.SimType)
	io.Uint8(&enumValue1)
	x.SimType = SimulationTypeType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
