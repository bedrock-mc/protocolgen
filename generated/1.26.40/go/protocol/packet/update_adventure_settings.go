// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateAdventureSettings struct {
	AdventureSettings protocol.AdventureSettings
}

// Marshal reads or writes UpdateAdventureSettings using its canonical wire layout.
func (x *UpdateAdventureSettings) Marshal(io protocol.IO) {
	x.AdventureSettings.Marshal(io)
}
