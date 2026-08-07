// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type EducationSettings struct {
	EducationLevelSettings protocol.EducationLevelSettings
}

// Marshal reads or writes EducationSettings using its canonical wire layout.
func (x *EducationSettings) Marshal(io protocol.IO) {
	x.EducationLevelSettings.Marshal(io)
}

// ID returns the protocol ID for EducationSettings.
func (*EducationSettings) ID() uint32 { return IDEducationSettings }
