// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateAbilities struct {
	Data protocol.SerializedAbilitiesData
}

// Marshal reads or writes UpdateAbilities using its canonical wire layout.
func (x *UpdateAbilities) Marshal(io protocol.IO) {
	x.Data.Marshal(io)
}

// ID returns the protocol ID for UpdateAbilities.
func (*UpdateAbilities) ID() uint32 { return IDUpdateAbilities }
