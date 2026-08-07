// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SimulationType is an in-progress packet. We currently do not know the use case.
type SimulationType struct {
	// SimulationType is the simulation type selected.
	SimType protocol.SimulationTypeEnum
}

// Marshal reads or writes SimulationType using its canonical wire layout.
func (x *SimulationType) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.SimType, io.Uint8)
}

// ID returns the protocol ID for SimulationType.
func (*SimulationType) ID() uint32 { return IDSimulationType }
