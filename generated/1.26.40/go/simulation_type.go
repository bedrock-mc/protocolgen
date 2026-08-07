// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SimulationType struct {
	SimType SimulationTypeType
}

// Marshal reads or writes SimulationType using its canonical wire layout.
func (x *SimulationType) Marshal(io IO) {
	IntegerFunc(&x.SimType, io.Uint8)
}
